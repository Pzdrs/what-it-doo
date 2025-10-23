package common

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
)

func SetAuthCookies(w http.ResponseWriter, session *model.UserSession, rememberMe bool) {
	cookie := &http.Cookie{
		Name:     internal.SESSION_COOKIE_NAME,
		Value:    session.Token,
		HttpOnly: true,
		Path:     "/",
	}

	authFlag := &http.Cookie{
		Name:  internal.IS_AUTHENTICATED_COOKIE_NAME,
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

func ClearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     internal.SESSION_COOKIE_NAME,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1, // ensures deletion in some browsers
	})

	http.SetCookie(w, &http.Cookie{
		Name:   internal.IS_AUTHENTICATED_COOKIE_NAME,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
