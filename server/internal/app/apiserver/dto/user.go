package dto

import (
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/domain/model"
)

type UserDetails struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Online   bool      `json:"online" validate:"required"`
	LastSeen time.Time `json:"last_seen" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Bio      string    `json:"bio"`
}

func ToUserDetails(u model.User) UserDetails {
	return UserDetails{
		ID:       u.ID,
		Name:     u.Name,
		Online:   u.Online,
		LastSeen: u.LastSeen,
		Email:    u.Email,
		Bio:      u.Bio,
	}
}
