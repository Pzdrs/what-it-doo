package apiserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB(getenv func(string) string) error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		getenv("DB_USER"), getenv("DB_PASSWORD"), getenv("DB_HOST"), getenv("DB_PORT"), getenv("DB_NAME"),
	)
	log.Printf("Connecting to database with DSN: %s", dsn)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
	return nil
}

func NewServer(getenv func(string) string) http.Handler {
	if err := initDB(getenv); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	addRoutes(
		r,
	)
	var handler http.Handler = r
	//handler = someMiddleware(handler)
	//handler = someMiddleware2(handler)
	//handler = someMiddleware3(handler)
	return handler
}
