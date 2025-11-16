package dto

import "pycrs.cz/what-it-doo/internal/apiserver/model"

func MapUserToUserDetails(u model.User) UserDetails {
	return UserDetails{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Bio:   u.Bio,
	}
}
