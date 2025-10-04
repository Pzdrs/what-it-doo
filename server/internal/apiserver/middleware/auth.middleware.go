package middleware

import (
	"context"
	"net/http"

	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/service"
)

type contextKey string

const (
	authCookie string     = "wid_session"
	userIDKey  contextKey = "session"
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
			cookie, err := r.Cookie(authCookie)
			if err != nil {
				http.Error(w, "authentication required", http.StatusUnauthorized)
				return
			}

			session, valid := authService.FindSession(cookie.Value)
			if !valid {
				http.Error(w, "invalid session", http.StatusUnauthorized)
				return
			}

			ctx := WithSession(r.Context(), session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireUnauthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(authCookie)
		if err == nil {
			http.Error(w, "Already authenticated", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
