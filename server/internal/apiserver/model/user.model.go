package model

import "time"

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Bio      string `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}