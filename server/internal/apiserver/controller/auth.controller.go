package controller

import (
	"encoding/json"
	"net/http"

	"pycrs.cz/what-it-do/internal/apiserver/middleware"
	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/problem"
	"pycrs.cz/what-it-do/internal/apiserver/service"
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
//	@Success		200		"Login successful"
//	@Failure		400		{object}	problem.ProblemDetails
//	@Failure		401		{object}	problem.ProblemDetails
//	@Param			request	body		controller.HandleLogin.loginRequest	true	"Login request"
//	@Router			/auth/login [post]
func (c *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	type loginRequest struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	} //	@name	LoginRequest

	decoder := json.NewDecoder(r.Body)
	var t loginRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	user, err := c.authService.GetUserByEmail(t.Email)
	if err != nil || !c.authService.AuthenticateUser(t.Email, t.Password) {
		problem.WriteProblemDetails(w, problem.ProblemDetails{
			Type:     "https://wid.pycrs.cz/probs/auth/incorrect-credentials",
			Title:    "Incorrect credentials",
			Status:   http.StatusUnauthorized,
			Detail:   "Incorrect email or password",
			Instance: "/auth/login",
		})
		return
	}

	session, err := c.authService.CreateSession(user.ID, "web", "unknown")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setAuthCookies(w, &session, t.RememberMe)
}

func setAuthCookies(w http.ResponseWriter, session *model.UserSession, rememberMe bool) {
	cookie := &http.Cookie{
		Name:     "wid_session",
		Value:    session.Token,
		HttpOnly: true,
		Path:     "/",
	}

	authFlag := &http.Cookie{
		Name:  "wid_is_authenticated",
		Value: "true",
		Path:  "/",
	}

	if rememberMe {
		cookie.Expires = session.ExpiresAt
		authFlag.Expires = session.ExpiresAt
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, authFlag)
}

// HandleRegister
//
//	@Summary		Register user
//	@Id				register
//	@Description	Register a new user with credentials
//	@Accept			json
//	@Tags			Authentication
//	@Success		200		"User registered"
//	@Failure		400		"Invalid input or user already exists"
//	@Param			request	body	controller.HandleRegister.registerRequest	true	"Register request"
//	@Router			/auth/register [post]
func (c *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	type registerRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} //	@name	RegisterRequest
	decoder := json.NewDecoder(r.Body)
	var t registerRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	u, err := c.authService.RegisterUser(
		model.User{
			Email: t.Email,
		}, t.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("User registered: " + u.Email))
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
	http.SetCookie(w, &http.Cookie{
		Name:     "wid_session",
		Value:    "",
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Path:     "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "wid_is_authenticated",
		Value:   "",
		Expires: session.ExpiresAt,
		Path:    "/",
	})
}
