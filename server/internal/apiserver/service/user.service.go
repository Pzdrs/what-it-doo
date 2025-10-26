package service

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/queries"
)

type UserService struct {
	repository *repository.UserRepository
	config     config.Configuration
}

func mapUserToModel(user queries.User, config config.Configuration) model.User {
	return model.User{
		ID:        user.ID,
		Name:      user.Name.String,
		Email:     user.Email,
		AvatarUrl: getAvatarUrl(user, config.Gravatar),
		Bio:       user.Bio.String,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}

func getAvatarUrl(user queries.User, config config.GravatarConfig) string {
	if user.AvatarUrl.String != "" {
		return user.AvatarUrl.String
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

func NewUserService(userRepo *repository.UserRepository, config config.Configuration) *UserService {
	return &UserService{repository: userRepo, config: config}
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*model.User, error) {
	return func() (*model.User, error) {
		user, err := s.repository.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		u := mapUserToModel(user, s.config)
		return &u, nil
	}()
}
