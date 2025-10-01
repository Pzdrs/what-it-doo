package repository

import (
	"pycrs.cz/what-it-do/internal/database"
)

type SessionRepository struct {
	q *database.Queries
}

func NewSessionRepository(q *database.Queries) *SessionRepository {
	return &SessionRepository{q: q}
}