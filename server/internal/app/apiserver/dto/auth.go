package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDetails struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Online   bool      `json:"online" validate:"required"`
	LastSeen time.Time `json:"last_seen" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Bio      string    `json:"bio"`
}

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me" validate:"required"`
}

type LoginResponse struct {
	User UserDetails `json:"user" validate:"required"`
}

type LogoutResponse struct {
	Success     bool   `json:"success" validate:"required"`
	RedirectUrl string `json:"redirect_url" validate:"required"`
}

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegistrationResponse struct {
	User UserDetails `json:"user" validate:"required"`
}
