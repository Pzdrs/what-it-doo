package controller

import (
	"errors"
	"net/http"

	"pycrs.cz/what-it-doo/internal/apiserver/common"
	"pycrs.cz/what-it-doo/internal/apiserver/dto"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// HandleLogin
//
//	@Summary		Authenticate user
//	@Id				login
//	@Description	Authenticate user with username and password
//	@Accept			json
//	@Tags			Authentication
//	@Success		200		{object}	dto.LoginResponse
//	@Failure		400		{object}	problem.ProblemDetails
//	@Failure		401		{object}	problem.ProblemDetails
//	@Param			request	body		dto.LoginRequest	true	"Login request"
//	@Router			/auth/login [post]
func (c *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	req, ok := common.DecodeAndValidate[dto.LoginRequest](w, r)
	if !ok {
		return
	}

	user, err := c.authService.GetUserByEmail(req.Email)
	if err != nil || !c.authService.AuthenticateUser(req.Email, req.Password) {
		problem.WriteProblemDetails(w, problem.NewProblemDetails(
			r, http.StatusUnauthorized,
			"Incorrect credentials",
			"Incorrect email or password",
			"auth/incorrect-credentials",
		))
		return
	}

	session, err := c.authService.CreateSession(user.ID, "web", "unknown")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SetAuthCookies(w, &session, req.RememberMe)
	common.WriteJSON(w, http.StatusCreated, dto.LoginResponse{
		User: dto.MapUserToUserDetails(user),
	})
}

// HandleRegister
//
//	@Summary		Register user
//	@Id				register
//	@Description	Register a new user with credentials
//	@Accept			json
//	@Tags			Authentication
//	@Success		201			{object}	dto.RegistrationResponse
//	@Failure		400			{object}	problem.ProblemDetails
//	@Failure		500			{object}	problem.ProblemDetails
//	@Param			request		body		dto.RegistrationRequest	true	"Register request"
//	@Param			autoLogin	query		bool					false	"Automatically log in the user after registration"
//	@Router			/auth/register [post]
func (c *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	req, ok := common.DecodeAndValidate[dto.RegistrationRequest](w, r)
	if !ok {
		return
	}

	user, err := c.authService.RegisterUser(
		model.User{
			Email: req.Email,
			Name:  req.Name,
		}, req.Password,
	)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			problem.WriteProblemDetails(w, problem.NewProblemDetails(
				r, http.StatusBadRequest,
				"User already exists",
				"A user with the given email already exists",
				"auth/email-taken",
			))
			return
		} else {
			problem.WriteProblemDetails(w, problem.NewInternalServerError(r, err))
			return
		}
	}

	if r.URL.Query().Get("autoLogin") == "true" {
		session, err := c.authService.CreateSession(user.ID, "web", "unknown")
		if err != nil {
			problem.WriteProblemDetails(w, problem.NewInternalServerError(r, err))
			return
		}
		common.SetAuthCookies(w, &session, true)
	}

	common.WriteJSON(w, http.StatusCreated, dto.RegistrationResponse{
		User: dto.MapUserToUserDetails(user),
	})
}

// HandleLogout
//
//	@Summary		Logout user
//	@Id				logout
//	@Description	Logout the authenticated user
//	@Tags			Authentication
//	@Success		200	{object}	dto.LogoutResponse
//	@Failure		401	{object}	problem.ProblemDetails
//	@Router			/auth/logout [post]
func (c *AuthController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.SessionFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}

	err := c.authService.Logout(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ClearAuthCookies(w)
}
