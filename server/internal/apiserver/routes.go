package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"pycrs.cz/what-it-do/internal/apiserver/controller"
	"pycrs.cz/what-it-do/internal/apiserver/service"
)

func addRoutes(
	r *chi.Mux,
	authService *service.AuthService,
	chatService *service.ChatService,
) {
	browserOnly := newBrowserOnly("This endpoint is only accessible from a web browser")

	authController := controller.NewAuthController(authService)
	chatController := controller.NewChatController(chatService)

	r.With(browserOnly).Get("/hello", controller.HandleHello)

	r.Route("/auth", func(r chi.Router) {
		r.With().Post("/login", authController.HandleLogin)
		r.With(requireAuthenticated).Post("/logout", authController.HandleLogout)
		r.With().Post("/register", authController.HandleRegister)
	})

	r.Route("/chats", func(r chi.Router) {
		r.Get("/", chatController.HandleAllChats)
	})

	r.Handle("/", http.NotFoundHandler())
}
