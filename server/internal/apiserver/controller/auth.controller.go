package controller

import (
	"encoding/json"
	"net/http"

	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (c *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t loginRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	user, err := c.authService.GetUserByUsername(t.Username)
	if err != nil || !c.authService.AuthenticateUser(t.Username, t.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	
	session, err := c.authService.CreateSession(user.ID, "web", "unknown")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "wid_session",
		Value: session.Token,
	})
}

func (c *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t registerRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	u, err := c.authService.RegisterUser(
		model.User{
			Username: t.Username, Email: t.Email,
		}, t.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("User registered: " + u.Username))
}

func (c *AuthController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout endpoint"))
}
