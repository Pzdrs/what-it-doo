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
	_, err := r.q.GetUserByUsername(context.Background(), username)
	return err == nil
}

func (r UserRepository) SaveUser(user database.User) (database.User, error) {
	fmt.Println("Saving user:", user)
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
	return r.q.GetUserByUsername(context.Background(), username)
}
