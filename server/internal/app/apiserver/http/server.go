package apiserver

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pycrs.cz/what-it-doo/internal/app/apiserver/presence"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/ws"
)

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

func NewServer(
	ctx context.Context, config config.Configuration,
	authService service.AuthService,
	chatService service.ChatService,
	userService service.UserService,
	sessionService service.SessionService,
	bus bus.CommnunicationBus,
	gatewayID string,
	connectionManager ws.ConnectionManager,
	presenceManager *presence.PresenceManager,
) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(api chi.Router) {
		addRoutes(
			ctx,
			api,
			gatewayID,
			authService,
			chatService,
			userService,
			sessionService,
			config,
			bus,
			connectionManager,
			presenceManager,
		)
	})

	// Static files (Svelte frontend)
	r.Get("/*", spaHandler("./web"))

	return r
}
