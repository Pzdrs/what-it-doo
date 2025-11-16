package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/bootstrap"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/queries"
	"pycrs.cz/what-it-doo/internal/worker"
)

func runWorker(ctx context.Context, q *queries.Queries, redisClient *redis.Client, config config.Configuration) error {
	userRepository := repository.NewUserRepository(q)
	chatRepository := repository.NewChatRepository(q)

	chatService := service.NewChatService(chatRepository, userRepository, config)

	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get hostname: %w", err)
	}

	for {
		resp, err := redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "workers",
			Consumer: hostname,
			Streams:  []string{"stream:tasks", ">"},
		}).Result()

		if ctx.Err() != nil {
			fmt.Println("Worker: context cancelled, exiting...")
			return nil
		}

		if err == redis.Nil {
			continue
		}
		if err != nil {
			return fmt.Errorf("worker redis read error: %w", err)
		}

		for _, msg := range resp[0].Messages {
			ttype, ok := msg.Values["type"].(string)
			if !ok {
				fmt.Println("❌ Invalid message format")
				redisClient.XAck(ctx, "stream:tasks", "workers", msg.ID)
				continue
			}

			switch ttype {
			case "message":
				var payload worker.MessagePayload
				if err := json.Unmarshal([]byte(msg.Values["payload"].(string)), &payload); err != nil {
					fmt.Println("❌ Failed to unmarshal message payload:", err)
					continue
				}

				if err := worker.ProcessMessageTask(ctx, chatService, payload); err != nil {
					fmt.Println("❌ Failed to process message task:", err)
					continue
				}
				fmt.Println("✅ Message task processed successfully")
			default:
				fmt.Println("⚠️ Unknown task type:", ttype)
			}

			redisClient.XAck(ctx, "stream:tasks", "workers", msg.ID)
			// i just need a queue bruh
			redisClient.XDel(ctx, "stream:tasks", msg.ID)
		}
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	config, err := bootstrap.InitConfig()
	if err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	connPool, err := bootstrap.InitDB(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer connPool.Close()

	redisClient, err := bootstrap.InitRedis(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to initialize redis: %w", err)
	}
	defer redisClient.Close()

	q := queries.New(connPool)

	bootstrap.EnsureStreamGroup(ctx, redisClient)

	// Worker goroutine
	errCh := make(chan error, 1)

	go func() {
		errCh <- runWorker(ctx, q, redisClient, config)
	}()

	fmt.Println("Worker started. Press Ctrl+C to stop.")

	// Wait for either:
	// - Ctrl+C (ctx.Done)
	// - worker error
	select {
	case <-ctx.Done():
		fmt.Println("Shutdown signal received.")
		// context cancels and unblocks XREADGROUP
		return nil

	case err := <-errCh:
		return fmt.Errorf("worker terminated with error: %w", err)
	}
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
