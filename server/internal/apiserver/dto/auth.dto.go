package dto

type LogoutResponse struct {
	Success     bool   `json:"success"`
	RedirectUrl string `json:"redirect_url"`
}
