package apiserver

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/repository"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/queries"
	"pycrs.cz/what-it-doo/internal/ws"
)

type Server struct {
	ctx               context.Context
	gatewayID         string
	connectionManager ws.ConnectionManager

	Handler http.Handler

	AuthService    service.AuthService
	ChatService    service.ChatService
	UserService    service.UserService
	SessionService service.SessionService
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

func NewServer(ctx context.Context, q *queries.Queries, config config.Configuration, bus bus.CommnunicationBus, gatewayID string, connectionManager ws.ConnectionManager) *Server {
	userRepository := repository.NewUserRepository(q)
	sessionRepository := repository.NewSessionRepository(q)
	chatRepository := repository.NewChatRepository(q)

	server := &Server{
		ctx:               ctx,
		gatewayID:         gatewayID,
		AuthService:       service.NewAuthService(userRepository, sessionRepository, config),
		ChatService:       service.NewChatService(chatRepository, userRepository, config),
		UserService:       service.NewUserService(userRepository, config),
		SessionService:    service.NewSessionService(userRepository, sessionRepository, config),
		connectionManager: connectionManager,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(api chi.Router) {
		addRoutes(
			ctx,
			api,
			gatewayID,
			server.AuthService,
			server.ChatService,
			server.UserService,
			server.SessionService,
			config,
			bus,
			connectionManager,
		)
	})

	// Static files (Svelte frontend)
	r.Get("/*", spaHandler("./web"))

	server.Handler = r

	return server
}
