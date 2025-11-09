package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/config"
)

type UserService interface {
	// GetUserByID retrieves a user by their ID.
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	// GetByEmail retrieves a user by their email.
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
	config     config.Configuration
}

func getAvatarUrl(user model.User, config config.GravatarConfig) string {
	if user.AvatarUrl != "" {
		return user.AvatarUrl
	}

	if config.Enabled {
		hash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(user.Email))))
		return strings.NewReplacer(
			"{{hash}}", hex.EncodeToString(hash[:]),
			"{{size}}", strconv.Itoa(80),
		).Replace(config.Url)
	}

	return ""
}

func NewUserService(userRepo repository.UserRepository, config config.Configuration) UserService {
	return &userService{repository: userRepo, config: config}
}

func (s *userService) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := s.repository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.AvatarUrl = getAvatarUrl(*user, s.config.Gravatar)
	return user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

var _ UserService = (*userService)(nil)
