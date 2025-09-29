package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
	"pycrs.cz/what-it-do/internal/apiserver/services"
	"pycrs.cz/what-it-do/internal/database"
)

type Server struct {
	Handler http.Handler

	authService *services.AuthService
}

func NewServer(q *database.Queries) *Server {
	userRepository := repository.NewUserRepository(q)
	
	server := &Server{
		authService: services.NewAuthService(userRepository),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	addRoutes(
		r,
		server.authService,
	)

	server.Handler = r

	return server
}
