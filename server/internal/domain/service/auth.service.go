package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"pycrs.cz/what-it-doo/internal/app/apiserver/common"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/model"
	"pycrs.cz/what-it-doo/internal/domain/repository"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type AuthService interface {
	// Register registers a new user with the given password.
	RegisterUser(ctx context.Context, user model.User, password string) (model.User, error)
	// AuthenticateUser checks if the provided email and password are valid.
	AuthenticateUser(ctx context.Context, email, password string) bool
	// LogoutUser logs out the user by revoking their session.
	LogoutUser(ctx context.Context, session model.UserSession) error
}

type authService struct {
	repository        repository.UserRepository
	sessionRepository repository.SessionRepository
	config            config.Configuration
}

func NewAuthService(repo repository.UserRepository, sessionRepo repository.SessionRepository, config config.Configuration) AuthService {
	return &authService{
		repository:        repo,
		sessionRepository: sessionRepo,
		config:            config,
	}
}

func (s *authService) RegisterUser(ctx context.Context, user model.User, password string) (model.User, error) {
	if s.repository.UserExists(ctx, user.Email) {
		return model.User{}, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	u := model.User{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: string(hashedPassword),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err = s.repository.Create(ctx, &u)
	if err != nil {
		return model.User{}, err
	}

	u.AvatarUrl = common.GetAvatarUrl(u, s.config.Gravatar)

	return u, nil
}

func (s *authService) AuthenticateUser(ctx context.Context, email, password string) bool {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func (s *authService) LogoutUser(ctx context.Context, session model.UserSession) error {
	return s.sessionRepository.DeleteByID(ctx, session.ID)
}

var _ AuthService = (*authService)(nil)
