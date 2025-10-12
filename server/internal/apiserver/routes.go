package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-do/api" // Swagger docs
	"pycrs.cz/what-it-do/internal/apiserver/controller"
	"pycrs.cz/what-it-do/internal/apiserver/middleware"
	"pycrs.cz/what-it-do/internal/apiserver/service"
)

func addRoutes(
	r chi.Router,
	authService *service.AuthService,
	chatService *service.ChatService,
	userService *service.UserService,
) {
	RequireAuthenticated := middleware.RequireAuthenticated(authService, userService)
	RequireUnauthenticated := middleware.RequireUnauthenticated(authService)

	authController := controller.NewAuthController(authService)
	chatController := controller.NewChatController(chatService)
	userController := controller.NewUserController(userService)

	r.Route("/auth", func(r chi.Router) {
		r.With().Post("/login", authController.HandleLogin)
		r.With(RequireAuthenticated).Post("/logout", authController.HandleLogout)
		r.With(RequireUnauthenticated).Post("/register", authController.HandleRegister)
	})

	r.Route("/users", func(r chi.Router) {
		r.With(RequireAuthenticated).Get("/me", userController.HandleGetMyself)
	})

	r.Route("/chats", func(r chi.Router) {
		r.Get("/", chatController.HandleAllChats)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Handle("/", http.NotFoundHandler())
}
