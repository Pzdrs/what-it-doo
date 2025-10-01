package repository

import (
	"context"
	"fmt"

	"pycrs.cz/what-it-do/internal/database"
)

type UserRepository struct {
	q *database.Queries
}

func NewUserRepository(q *database.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) UserExists(username string) bool {
	_, err := r.q.GetUserByUsernameOrEmail(context.Background(), username)
	return err == nil
}

func (r UserRepository) SaveUser(user database.User) (database.User, error) {
	if user.Email == "" || user.Username == "" {
		return database.User{}, fmt.Errorf("username and email are required")
	}

	user, err := r.q.CreateUser(context.Background(), database.CreateUserParams{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		AvatarUrl:      user.AvatarUrl,
		Bio:            user.Bio,
	})
	return user, err
}

func (r *UserRepository) GetUserByUsername(username string) (database.User, error) {
	return r.q.GetUserByUsernameOrEmail(context.Background(), username)
}

func (r *UserRepository) CreateSession(params database.CreateSessionParams) (database.Session, error) {
	return r.q.CreateSession(context.Background(), params)
}
