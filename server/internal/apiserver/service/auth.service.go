package service

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
	"pycrs.cz/what-it-do/internal/database"
)

type AuthService struct {
	repository *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

func mapUserToModel(user database.User) model.User {
	return model.User{
		ID:        user.ID,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl.String,
		Bio:       user.Bio.String,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}

func (s *AuthService) RegisterUser(user model.User, password string) (model.User, error) {
	if s.repository.UserExists(user.Username) {
		return model.User{}, fmt.Errorf("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u, err := s.repository.SaveUser(database.User{
		Username:       user.Username,
		FirstName:      pgtype.Text{String: user.FirstName, Valid: true},
		LastName:       pgtype.Text{String: user.LastName, Valid: true},
		Email:          user.Email,
		HashedPassword: pgtype.Text{String: string(hashedPassword), Valid: true},
		CreatedAt:      pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now()},
	})
	if err != nil {
		return model.User{}, err
	}
	return mapUserToModel(u), nil
}

func (s *AuthService) AuthenticateUser(username, password string) bool {
	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword.String), []byte(password))
	return err == nil
}