package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"pycrs.cz/what-it-do/internal/apiserver/services"
)

func addRoutes(
	r *chi.Mux,
	authService *services.AuthService,
) {
	browserOnly := newBrowserOnly("This endpoint is only accessible from a web browser")

	authController := NewAuthController(authService)

	r.With(browserOnly).Get("/hello", handleHello)

	r.Route("/auth", func(r chi.Router) {
		r.With(requireUnauthenticated).Post("/login", authController.handleLogin)
		r.With(requireAuthenticated).Post("/logout", authController.handleLogout)
		r.With(requireUnauthenticated).Post("/register", authController.handleRegister)
	})

	r.Handle("/", http.NotFoundHandler())
}
