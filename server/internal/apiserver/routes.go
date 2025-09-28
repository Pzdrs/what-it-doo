package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func addRoutes(
	r *chi.Mux,
) {
	browserOnly := newBrowserOnly("This endpoint is only accessible from a web browser")

	r.With(browserOnly).Get("/hello", handleHello)

	r.Route("/auth", func(r chi.Router) {
		r.With(requireUnauthenticated).Post("/login", handleLogin)
		r.With(requireAuthenticated).Post("/logout", handleLogout)
		r.With(requireUnauthenticated).Post("/register", handleRegister)
	})

	r.Handle("/", http.NotFoundHandler())
}
