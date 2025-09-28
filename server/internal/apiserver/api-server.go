package apiserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewServer() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	addRoutes(
		r,
	)
	var handler http.Handler = r
	//handler = someMiddleware(handler)
	//handler = someMiddleware2(handler)
	//handler = someMiddleware3(handler)
	return handler
}
