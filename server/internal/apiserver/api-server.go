package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
	"pycrs.cz/what-it-do/internal/apiserver/service"
	"pycrs.cz/what-it-do/internal/database"
)

type Server struct {
	Handler http.Handler

	authService *service.AuthService
	chatService *service.ChatService
}

func NewServer(q *database.Queries) *Server {
	userRepository := repository.NewUserRepository(q)
	chatRepository := repository.NewChatRepository(q)
	
	server := &Server{
		authService: service.NewAuthService(userRepository),
		chatService: service.NewChatService(chatRepository),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	addRoutes(
		r,
		server.authService,
		server.chatService,
	)

	server.Handler = r

	return server
}
