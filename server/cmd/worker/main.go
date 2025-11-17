package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/app/worker"
	"pycrs.cz/what-it-doo/internal/bootstrap"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/bus/payload"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/queries"
)

func runWorker(ctx context.Context, q *queries.Queries, bus bus.CommnunicationBus, config config.Configuration) error {
	userRepository := repository.NewUserRepository(q)
	chatRepository := repository.NewChatRepository(q)

	chatService := service.NewChatService(chatRepository, userRepository, config)

	taskChan, err := bus.ConsumeTasks(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case task := <-taskChan:
			switch task.Type {
			case "message":
				var payload payload.MessageTaskPayload
				if err := json.Unmarshal(task.Payload, &payload); err != nil {
					fmt.Println("❌ Failed to unmarshal message payload:", err)
					continue
				}

				if err := worker.ProcessMessageTask(ctx, chatService, bus, payload); err != nil {
					fmt.Println("❌ Failed to process message task:", err)
					continue
				}
				fmt.Println("✅ Message task processed successfully")
			default:
				fmt.Println("⚠️ Unknown task type:", task.Type)
			}

			bus.AckTask(ctx, task.ID)
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

	bus := bus.NewRedisCommunicationBus(redisClient)

	bootstrap.EnsureStreamGroup(ctx, redisClient)

	// Worker goroutine
	errCh := make(chan error, 1)

	go func() {
		errCh <- runWorker(ctx, q, bus, config)
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
