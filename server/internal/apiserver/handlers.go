package apiserver

import (
	"encoding/json"
	"net/http"

	"pycrs.cz/what-it-do/internal/apiserver/services"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t loginRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	if services.AuthenticateUser(t.Username, t.Password) {
		w.Write([]byte("Login successful"))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t registerRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	msg, err := services.RegisterUser(t.Username, t.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Write([]byte(msg))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout endpoint"))
}
