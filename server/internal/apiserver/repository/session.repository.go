package repository

import (
	"context"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/queries"
)

type SessionRepository struct {
	q *queries.Queries
}

func NewSessionRepository(q *queries.Queries) *SessionRepository {
	return &SessionRepository{q: q}
}

func (r *SessionRepository) GetSessionByToken(token string) (queries.Session, error) {
	return r.q.GetSessionByToken(context.Background(), token)
}

func (r *SessionRepository) CreateSession(params queries.CreateSessionParams) (queries.Session, error) {
	return r.q.CreateSession(context.Background(), params)
}

func (r *SessionRepository) DeleteSessionByID(sessionID uuid.UUID) error {
	return r.q.DeleteSessionByID(context.Background(), sessionID)
}
