package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"pycrs.cz/what-it-do/internal/models"
)

var DB *sql.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "root"
	PASSWORD = "root"
	DBNAME   = "whatitdo"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]models.User, 0)

	rows, err := DB.Query(`SELECT id, first_name, last_name, email FROM users`)
	if err != nil {
		log.Println("failed to execute query:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			log.Println("failed to scan:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	var err error
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)

	DB, err = sql.Open("postgres", connString) // note the plain '='
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	userR := r.PathPrefix("/users").Subrouter()
	userR.HandleFunc("", GetAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
