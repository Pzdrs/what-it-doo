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
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"pycrs.cz/what-it-doo/internal/apiserver"
	"pycrs.cz/what-it-doo/internal/database"
	"pycrs.cz/what-it-doo/internal/validation"
	"pycrs.cz/what-it-doo/pkg/version"
)

func initDB(config *apiserver.Configuration) (*pgx.Conn, error) {
	var dsn string

	if config.Database.URL != "" {
		dsn = config.Database.URL
	} else {
		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
	}

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to database")
	return conn, nil
}

func initConfig() (apiserver.Configuration, error) {
	config := viper.NewWithOptions(viper.ExperimentalBindStruct())
	config.AutomaticEnv()
	config.SetConfigName("wid")
	config.SetEnvPrefix("WID")

	config.AddConfigPath(".")
	config.AddConfigPath("/etc/whatitdoo")

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config.SetDefault("server.port", 8080)
	config.SetDefault("database.port", 5432)

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return apiserver.Configuration{}, fmt.Errorf("failed to read config: %w", err)
		}
		log.Println("No configuration file found")
	} else {
		log.Println("Using configuration file:", config.ConfigFileUsed())
	}

	var cfg apiserver.Configuration
	if err := config.Unmarshal(&cfg); err != nil {
		return apiserver.Configuration{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterStructValidation(
		validation.DbConfigStructLevelValidation,
		apiserver.DatabaseConfiguration{},
	)
	if err := validate.Struct(cfg); err != nil {
		return apiserver.Configuration{}, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

// @title			What-it-doo API
// @version		1.0
// @description	API for the messanger of the future - What-it-doo.
// @BasePath		/api/v1
func run(ctx context.Context, getenv func(string) string, w io.Writer, args []string) error {
	log.Printf("Starting what-it-doo server version %s\n", version.Version)

	config, err := initConfig()
	if err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	godotenv.Load()
	conn, err := initDB(&config)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	q := database.New(conn)

	server := apiserver.NewServer(q)
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
