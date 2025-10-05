package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"pycrs.cz/what-it-do/internal/database"
)

type UserRepository struct {
	q *database.Queries
}

func NewUserRepository(q *database.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) UserExists(username string) bool {
	_, err := r.q.GetUserByEmail(context.Background(), username)
	return err == nil
}

func (r UserRepository) SaveUser(user database.User) (database.User, error) {
	if user.Email == "" {
		return database.User{}, fmt.Errorf("email is required")
	}

	user, err := r.q.CreateUser(context.Background(), database.CreateUserParams{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		AvatarUrl:      user.AvatarUrl,
		Bio:            user.Bio,
	})
	return user, err
}

func (r *UserRepository) GetUserByEmail(email string) (database.User, error) {
	return r.q.GetUserByEmail(context.Background(), email)
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (*database.User, error) {
	user, err := r.q.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
