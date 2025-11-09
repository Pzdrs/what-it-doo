package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/queries"
)

type SessionRepository interface {
	Create(ctx context.Context, session *model.UserSession) error
	GetByToken(ctx context.Context, token string) (*model.UserSession, error)
	DeleteByID(ctx context.Context, sessionID uuid.UUID) error
}

type postgresSessionRepository struct {
	q *queries.Queries
}

func NewSessionRepository(q *queries.Queries) SessionRepository {
	return &postgresSessionRepository{q: q}
}

func (r *postgresSessionRepository) GetByToken(ctx context.Context, token string) (*model.UserSession, error) {
	s, err := r.q.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	session := &model.UserSession{
		ID:         s.ID,
		UserID:     s.UserID,
		Token:      s.Token,
		DeviceType: s.DeviceType.String,
		DeviceOs:   s.DeviceOs.String,
		CreatedAt:  s.CreatedAt.Time,
		ExpiresAt:  s.ExpiresAt.Time,
	}

	if s.RevokedAt.Valid {
		session.RevokedAt = &s.RevokedAt.Time
	}

	return session, nil
}

func (r *postgresSessionRepository) Create(ctx context.Context, session *model.UserSession) error {
	_, err := r.q.CreateSession(ctx, queries.CreateSessionParams{
		UserID:     session.UserID,
		Token:      session.Token,
		DeviceType: pgtype.Text{String: session.DeviceType, Valid: true},
		DeviceOs:   pgtype.Text{String: session.DeviceOs, Valid: true},
		ExpiresAt:  pgtype.Timestamptz{Time: session.ExpiresAt, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresSessionRepository) DeleteByID(ctx context.Context, sessionID uuid.UUID) error {
	return r.q.DeleteSessionByID(ctx, sessionID)
}

var _ SessionRepository = (*postgresSessionRepository)(nil)
