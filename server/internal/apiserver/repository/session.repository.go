package repository

import (
	"context"

	"pycrs.cz/what-it-do/internal/database"
)

type SessionRepository struct {
	q *database.Queries
}

func NewSessionRepository(q *database.Queries) *SessionRepository {
	return &SessionRepository{q: q}
}

func (r *SessionRepository) GetSessionByToken(token string) (database.Session, error) {
	return r.q.GetSessionByToken(context.Background(), token)
}

func (r *SessionRepository) CreateSession(params database.CreateSessionParams) (database.Session, error) {
	return r.q.CreateSession(context.Background(), params)
}