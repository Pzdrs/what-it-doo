package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"pycrs.cz/what-it-doo/internal/apiserver"
	"pycrs.cz/what-it-doo/internal/bootstrap"
	"pycrs.cz/what-it-doo/internal/queries"
	"pycrs.cz/what-it-doo/pkg/version"
)

// @title			What-it-doo API
// @version		1.0
// @description	API for the messanger of the future - What-it-doo.
// @BasePath		/api/v1
func run(ctx context.Context, getenv func(string) string, w io.Writer, args []string) error {
	log.Printf("Starting what-it-doo server version %s\n", version.Version)

	config, err := bootstrap.InitConfig()
	if err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	conn, err := bootstrap.InitDB(&config)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	redisClient, err := bootstrap.InitRedis(&config)
	if err != nil {
		return fmt.Errorf("failed to initialize redis: %w", err)
	}
	q := queries.New(conn)

	server := apiserver.NewServer(q, redisClient)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", strconv.Itoa(config.Server.Port)),
		Handler: server.Handler,
	}

	log.Printf("Listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("error listening and serving: %s\n", err)
	}
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Getenv, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
