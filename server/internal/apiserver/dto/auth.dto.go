package dto

import "github.com/google/uuid"

type UserDetails struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Bio   string    `json:"bio"`
}

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginResponse struct {
	User UserDetails `json:"user" validate:"required"`
}

type LogoutResponse struct {
	Success     bool   `json:"success"`
	RedirectUrl string `json:"redirect_url"`
}

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegistrationResponse struct {
	User UserDetails `json:"user" validate:"required"`
}
