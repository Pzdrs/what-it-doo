package controller

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal/app/apiserver/common"
	"pycrs.cz/what-it-doo/internal/app/apiserver/dto"
	"pycrs.cz/what-it-doo/internal/app/apiserver/http/middleware"
	"pycrs.cz/what-it-doo/internal/domain/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// HandleGetMyself
//
//	@Summary		Get current user
//	@Description	Get details of the currently authenticated user
//	@Id				getMyself
//	@Tags			users
//	@Produce		json
//	@Router			/users/me [get]
func (c *UserController) HandleGetMyself(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.SessionFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}

	user, err := c.userService.GetByID(r.Context(), session.UserID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	common.Encode(w, r, 200, dto.ToUserDetails(*user))
}
