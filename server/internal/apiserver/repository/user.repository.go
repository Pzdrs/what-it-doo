package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/queries"
)

type UserRepository struct {
	q *queries.Queries
}

func NewUserRepository(q *queries.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) UserExists(username string) bool {
	_, err := r.q.GetUserByEmail(context.Background(), username)
	return err == nil
}

func (r UserRepository) SaveUser(user queries.User) (queries.User, error) {
	if user.Email == "" {
		return queries.User{}, fmt.Errorf("email is required")
	}

	user, err := r.q.CreateUser(context.Background(), queries.CreateUserParams{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		AvatarUrl:      user.AvatarUrl,
		Bio:            user.Bio,
	})
	return user, err
}

func (r *UserRepository) GetUserByEmail(email string) (queries.User, error) {
	return r.q.GetUserByEmail(context.Background(), email)
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (*queries.User, error) {
	user, err := r.q.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
