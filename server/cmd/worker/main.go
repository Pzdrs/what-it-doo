package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"pycrs.cz/what-it-doo/internal/app/worker"
	"pycrs.cz/what-it-doo/internal/bootstrap"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/queries"
)

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
		errCh <- worker.RunWorker(ctx, q, bus, config)
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
