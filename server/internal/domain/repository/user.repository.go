package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"pycrs.cz/what-it-doo/internal/domain/model"
	"pycrs.cz/what-it-doo/internal/queries"
)

type UserRepository interface {
	UserExists(ctx context.Context, email string) bool
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)

	SetOnline(ctx context.Context, userID uuid.UUID) error
	SetOffline(ctx context.Context, userID uuid.UUID, lastSeen time.Time) error
}

type postgresUserRepository struct {
	q *queries.Queries
}

func NewUserRepository(q *queries.Queries) UserRepository {
	return &postgresUserRepository{q: q}
}

func (r *postgresUserRepository) UserExists(ctx context.Context, email string) bool {
	_, err := r.q.GetUserByEmail(ctx, email)
	return err == nil
}

func (r *postgresUserRepository) Create(ctx context.Context, user *model.User) error {
	createdUser, err := r.q.CreateUser(ctx, queries.CreateUserParams{
		Name:           pgtype.Text{String: user.Name, Valid: true},
		Email:          user.Email,
		HashedPassword: pgtype.Text{String: user.HashedPassword, Valid: true},
		Bio:            pgtype.Text{String: user.Bio, Valid: true},
	})

	if err != nil {
		return err
	}

	user.ID = createdUser.ID
	user.CreatedAt = createdUser.CreatedAt.Time
	user.UpdatedAt = createdUser.UpdatedAt.Time

	return nil
}

func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	userModel := mapUser(user)
	return &userModel, nil
}

func (r *postgresUserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := r.q.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	userModel := mapUser(user)
	return &userModel, nil
}

func (r *postgresUserRepository) SetOnline(ctx context.Context, userID uuid.UUID) error {
	return r.q.SetUserOnline(ctx, userID)
}

func (r *postgresUserRepository) SetOffline(ctx context.Context, userID uuid.UUID, lastSeen time.Time) error {
	return r.q.SetUserOffline(ctx, queries.SetUserOfflineParams{
		ID:           userID,
		LastActiveAt: pgtype.Timestamptz{Time: lastSeen, Valid: true},
	})
}

func mapUser(user queries.User) model.User {
	return model.User{
		ID:             user.ID,
		Name:           user.Name.String,
		Email:          user.Email,
		HashedPassword: user.HashedPassword.String,
		Bio:            user.Bio.String,
		Online:         user.IsOnline.Bool,
		LastSeen:       user.LastActiveAt.Time,
		CreatedAt:      user.CreatedAt.Time,
		UpdatedAt:      user.UpdatedAt.Time,
	}
}

// Compile-time check
var _ UserRepository = (*postgresUserRepository)(nil)
