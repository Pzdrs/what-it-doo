package apiserver

import (
	"net/http"
	"strings"

	"pycrs.cz/what-it-do/internal/apiserver/service"
)

const (
	authCookie = "wid_session"
)

func newBrowserOnly(errorMsg string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.Header.Get("User-Agent"), "Mozilla") {
				http.Error(w, errorMsg, http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func requireAuthenticated(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(authCookie)
			if err != nil {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			} else {
				if !authService.ValidateSession(cookie.Value) {
					http.Error(w, "Invalid session", http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func requireUnauthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(authCookie)
		if err == nil {
			http.Error(w, "Already authenticated", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
