package middleware

import (
	"context"

	"pycrs.cz/what-it-doo/internal/domain/model"
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
