package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "pycrs.cz/what-it-do/api" // Swagger docs
	"pycrs.cz/what-it-do/internal/apiserver/controller"
	"pycrs.cz/what-it-do/internal/apiserver/service"
)

func addRoutes(
	r *chi.Mux,
	authService *service.AuthService,
	chatService *service.ChatService,
) {
	browserOnly := newBrowserOnly("This endpoint is only accessible from a web browser")
	requireAuthenticated := requireAuthenticated(authService)

	authController := controller.NewAuthController(authService)
	chatController := controller.NewChatController(chatService)

	r.With(browserOnly).Get("/hello", controller.HandleHello)

	r.Route("/auth", func(r chi.Router) {
		r.With(requireUnauthenticated).Post("/login", authController.HandleLogin)
		r.With(requireAuthenticated).Post("/logout", authController.HandleLogout)
		r.With(requireUnauthenticated).Post("/register", authController.HandleRegister)
	})

	r.Route("/chats", func(r chi.Router) {
		r.Get("/", chatController.HandleAllChats)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Handle("/", http.NotFoundHandler())
}
