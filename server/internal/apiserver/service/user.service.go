package service

import (
	"github.com/google/uuid"
	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
	"pycrs.cz/what-it-do/internal/database"
)

type UserService struct {
	repository *repository.UserRepository
}

func mapUserToModel(user database.User) model.User {
	return model.User{
		ID:        user.ID,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl.String,
		Bio:       user.Bio.String,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{repository: userRepo}
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*model.User, error) {
	return func() (*model.User, error) {
		user, err := s.repository.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		u := mapUserToModel(*user)
		return &u, nil
	}()
}
