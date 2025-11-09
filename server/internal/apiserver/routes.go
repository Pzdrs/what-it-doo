package apiserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-doo/api" // Swagger docs
	"pycrs.cz/what-it-doo/internal/apiserver/controller"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
	"pycrs.cz/what-it-doo/internal/config"
)

func addRoutes(
	r chi.Router,
	authService service.AuthService,
	chatService service.ChatService,
	userService service.UserService,
	sessionService service.SessionService,
	config config.Configuration,
) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connectionManager := ws.NewConnectionManager()

	RequireAuthenticated := middleware.RequireAuthenticated(sessionService)
	RequireUnauthenticated := middleware.RequireUnauthenticated(sessionService)

	authController := controller.NewAuthController(authService, userService, sessionService)
	chatController := controller.NewChatController(chatService)
	userController := controller.NewUserController(userService)
	serverController := controller.NewServerController()
	socketController := controller.NewSocketController(upgrader, connectionManager)

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
