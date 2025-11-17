package apiserver

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-doo/api" // Swagger docs
	"pycrs.cz/what-it-doo/internal/apiserver/controller"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/config"
)

func addRoutes(
	ctx context.Context,
	r chi.Router,
	gatewayID string,
	authService service.AuthService,
	chatService service.ChatService,
	userService service.UserService,
	sessionService service.SessionService,
	config config.Configuration,
	bus bus.CommnunicationBus,
	connectionManager ws.ConnectionManager,
) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	RequireAuthenticated := middleware.RequireAuthenticated(sessionService)
	RequireUnauthenticated := middleware.RequireUnauthenticated(sessionService)

	authController := controller.NewAuthController(authService, userService, sessionService)
	chatController := controller.NewChatController(chatService, config)
	userController := controller.NewUserController(userService)
	serverController := controller.NewServerController()
	socketController := controller.NewSocketController(ctx, upgrader, connectionManager, bus, userService, gatewayID)

	r.Route("/server", func(r chi.Router) {
		r.Get("/about", serverController.HandleAbout)
		r.Get("/config", serverController.HandleConfig)
	})

	r.Route("/auth", func(r chi.Router) {
		r.With().Post("/login", authController.HandleLogin)
		r.With(RequireAuthenticated).Post("/logout", authController.HandleLogout)
		r.With(RequireUnauthenticated).Post("/register", authController.HandleRegister)
	})

	r.Route("/users", func(r chi.Router) {
		r.With(RequireAuthenticated).Get("/me", userController.HandleGetMyself)
	})

	r.With(RequireAuthenticated).Route("/chats", func(r chi.Router) {
		r.Get("/", chatController.HandleMyChats)
		r.Post("/", chatController.HandleCreateChat)
		r.Route("/{chat_id}", func(r chi.Router) {
			r.Get("/", chatController.HandleGetChat)
			r.Get("/messages", chatController.HandleGetChatMessages)
		})
	})

	r.With(RequireAuthenticated).Get("/ws", socketController.HandleWebSocket)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
