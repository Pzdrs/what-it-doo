package dto

import "github.com/google/uuid"

type UserDetails struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	Bio       string    `json:"bio"`
}

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginResponse struct {
	User UserDetails `json:"user"`
}

type LogoutResponse struct {
	Success     bool   `json:"success"`
	RedirectUrl string `json:"redirect_url"`
}

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegistrationResponse struct {
	Success bool        `json:"success"`
	User    UserDetails `json:"user"`
}
