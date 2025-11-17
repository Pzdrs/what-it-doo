package middleware

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal/app/apiserver"
	"pycrs.cz/what-it-doo/internal/app/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/domain/service"
)

func RequireAuthenticated(sessionService service.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(apiserver.SESSION_COOKIE_NAME)
			if err != nil {
				problem.Write(w, problem.New(
					r, http.StatusUnauthorized,
					"Unauthenticated",
					"Authentication is required to access this resource",
					"auth/unauthenticated",
				))
				return
			}

			session, valid := sessionService.GetByToken(r.Context(), cookie.Value)
			if !valid {
				problem.Write(w, problem.New(
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

func RequireUnauthenticated(sessionService service.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(apiserver.SESSION_COOKIE_NAME)
			if err == nil {
				_, valid := sessionService.GetByToken(r.Context(), cookie.Value)
				if valid {
					problem.Write(w, problem.New(
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
