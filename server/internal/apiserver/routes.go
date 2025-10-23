package apiserver

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-doo/api" // Swagger docs
	"pycrs.cz/what-it-doo/internal/apiserver/controller"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
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
	serverController := controller.NewServerController()

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

	r.Route("/chats", func(r chi.Router) {
		r.Get("/", chatController.HandleAllChats)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
