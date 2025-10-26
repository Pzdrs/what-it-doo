package controller

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal/apiserver/common"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// HandleGetMyself
//
//	@Summary		Get current user
//	@Description	Get details of the currently authenticated user
//	@Id				getMyself
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	model.User
//	@Router			/users/me [get]
func (c *UserController) HandleGetMyself(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.SessionFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}

	user, err := c.userService.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	common.WriteJSON(w, 200, user)
}
