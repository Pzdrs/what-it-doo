package apiserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-doo/api" // Swagger docs
	"pycrs.cz/what-it-doo/internal/apiserver/controller"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/config"
)

func addRoutes(
	r chi.Router,
	authService *service.AuthService,
	chatService service.ChatService,
	userService *service.UserService,
	config config.Configuration,
) {
	RequireAuthenticated := middleware.RequireAuthenticated(authService, userService)
	RequireUnauthenticated := middleware.RequireUnauthenticated(authService)

	authController := controller.NewAuthController(authService)
	chatController := controller.NewChatController(chatService)
	userController := controller.NewUserController(userService)
	serverController := controller.NewServerController()

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

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
		r.Route("/{chat_id}", func(r chi.Router) {
			r.Get("/", chatController.HandleGetChat)
			r.Get("/messages", chatController.HandleGetChatMessages)
		})
	})

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
			return
		}

		go func() {
			defer conn.Close()
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}
				// Handle incoming WebSocket messages
				fmt.Println("Received WebSocket message:", string(msg))
			}
		}()
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
