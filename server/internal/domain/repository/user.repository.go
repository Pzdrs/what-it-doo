package repository

import (
	"context"

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
	return &model.User{
		ID:             user.ID,
		Name:           user.Name.String,
		Email:          user.Email,
		HashedPassword: user.HashedPassword.String,
		Bio:            user.Bio.String,
		CreatedAt:      user.CreatedAt.Time,
		UpdatedAt:      user.UpdatedAt.Time,
	}, nil
}

func (r *postgresUserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := r.q.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:             user.ID,
		Name:           user.Name.String,
		Email:          user.Email,
		HashedPassword: user.HashedPassword.String,
		Bio:            user.Bio.String,
		CreatedAt:      user.CreatedAt.Time,
		UpdatedAt:      user.UpdatedAt.Time,
	}, nil
}

// Compile-time check
var _ UserRepository = (*postgresUserRepository)(nil)
