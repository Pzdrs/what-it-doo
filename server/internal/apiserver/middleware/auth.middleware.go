package middleware

import (
	"context"
	"net/http"

	"pycrs.cz/what-it-doo/internal"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

type contextKey string

const (
	userIDKey contextKey = "session"
)

func WithSession(ctx context.Context, session model.UserSession) context.Context {
	return context.WithValue(ctx, userIDKey, session)
}

func SessionFromContext(ctx context.Context) (model.UserSession, bool) {
	session, ok := ctx.Value(userIDKey).(model.UserSession)
	return session, ok
}

func RequireAuthenticated(authService *service.AuthService, userService *service.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(internal.SESSION_COOKIE_NAME)
			if err != nil {
				problem.WriteProblemDetails(w, problem.NewProblemDetails(
					r, http.StatusUnauthorized,
					"Unauthenticated",
					"Authentication is required to access this resource",
					"auth/unauthenticated",
				))
				return
			}

			session, valid := authService.FindSession(cookie.Value)
			if !valid {
				problem.WriteProblemDetails(w, problem.NewProblemDetails(
					r, http.StatusUnauthorized,
					"Invalid session",
					"The provided session either does not exist, is expired or has been revoked",
					"auth/invalid-session",
				))
				return
			}

			ctx := WithSession(r.Context(), session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireUnauthenticated(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(internal.SESSION_COOKIE_NAME)
			if err == nil {
				_, valid := authService.FindSession(cookie.Value)
				if valid {
					problem.WriteProblemDetails(w, problem.NewProblemDetails(
						r, http.StatusBadRequest,
						"Already authenticated",
						"This resource is only available for unauthenticated users",
						"auth/already-authenticated",
					))
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
