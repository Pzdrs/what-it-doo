package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"pycrs.cz/what-it-doo/internal/apiserver"
)

func InitDB(config *apiserver.Configuration) (*pgx.Conn, error) {
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
