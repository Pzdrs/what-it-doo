package apiserver

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var users = map[string]string{
	"alice": "password123",
	"bob":   "securepassword",
}

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
	if storedPassword, ok := users[t.Username]; ok {
		if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(t.Password)); err == nil {
			w.Write([]byte("Login successful"))
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t registerRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	if _, ok := users[t.Username]; ok {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	users[t.Username] = string(hashedPassword)
	w.Write([]byte("Registration successful"))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout endpoint"))
}
