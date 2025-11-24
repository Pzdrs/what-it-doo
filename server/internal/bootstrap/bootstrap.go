package bootstrap

import (
	"os"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/repository"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/queries"
)

func GenerateGatewayID() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "gateway-"+uuid.NewString()
	}
	return hostname
}

func InitServices(q *queries.Queries, config config.Configuration) (
	service.AuthService,
	service.ChatService,
	service.UserService,
	service.SessionService,
) {
	userRepository := repository.NewUserRepository(q)
	sessionRepository := repository.NewSessionRepository(q)
	chatRepository := repository.NewChatRepository(q)

	authService := service.NewAuthService(userRepository, sessionRepository, config)
	chatService := service.NewChatService(chatRepository, userRepository, config)
	userService := service.NewUserService(userRepository, config)
	sessionService := service.NewSessionService(userRepository, sessionRepository, config)

	return authService, chatService, userService, sessionService
}
