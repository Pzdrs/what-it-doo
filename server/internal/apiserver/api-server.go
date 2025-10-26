package apiserver

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/queries"
)

type Server struct {
	Handler http.Handler

	authService *service.AuthService
	chatService *service.ChatService
	userService *service.UserService
}

func spaHandler(staticDir string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(staticDir))

	return func(w http.ResponseWriter, r *http.Request) {
		// Does the path look like a file? (.js, .css, .png, etc.)
		if strings.Contains(filepath.Base(r.URL.Path), ".") {
			// Let FileServer try to serve it
			fs.ServeHTTP(w, r)
			return
		}

		// Otherwise, treat it as an SPA route → serve index.html
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	}
}

func NewServer(q *queries.Queries) *Server {
	userRepository := repository.NewUserRepository(q)
	sessionRepository := repository.NewSessionRepository(q)
	chatRepository := repository.NewChatRepository(q)

	server := &Server{
		authService: service.NewAuthService(userRepository, sessionRepository),
		chatService: service.NewChatService(chatRepository),
		userService: service.NewUserService(userRepository),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(api chi.Router) {
		addRoutes(
			api,
			server.authService,
			server.chatService,
			server.userService,
		)
	})

	// Static files (Svelte frontend)
	r.Get("/*", spaHandler("./static"))

	server.Handler = r

	return server
}
