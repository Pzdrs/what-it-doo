package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

type UserIdPLink struct {
	UserID      uuid.UUID
	ProviderUID string
}
