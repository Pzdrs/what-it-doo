package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/app/apiserver/common"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/model"
	"pycrs.cz/what-it-doo/internal/domain/repository"
)

type UserService interface {
	// GetUserByID retrieves a user by their ID.
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	// GetByEmail retrieves a user by their email.
	GetByEmail(ctx context.Context, email string) (model.User, error)
	SetPresence(ctx context.Context, userID uuid.UUID, online bool) error
}

type userService struct {
	repository repository.UserRepository
	config     config.Configuration
}

func NewUserService(userRepo repository.UserRepository, config config.Configuration) UserService {
	return &userService{repository: userRepo, config: config}
}

func (s *userService) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := s.repository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.AvatarUrl = common.GetAvatarUrl(*user, s.config.Gravatar)
	return user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func (s *userService) SetPresence(ctx context.Context, userID uuid.UUID, online bool) error {
	if online {
		return s.repository.SetOnline(ctx, userID)
	} else {
		return s.repository.SetOffline(ctx, userID, time.Now())
	}
}

var _ UserService = (*userService)(nil)
