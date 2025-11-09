package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/queries"
)

type ChatRepository interface {
	GetAll(ctx context.Context) ([]model.Chat, error)
	GetForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error)
	GetByID(ctx context.Context, chatID int64) (*model.Chat, error)
	GetMessagesForChat(ctx context.Context, chatID int64, limit int32, beforeTime time.Time) ([]model.Message, error)
	Create(ctx context.Context) (model.Chat, error)
	AddParticipant(ctx context.Context, chatID int64, userID uuid.UUID) error
	GetParticipants(ctx context.Context, chatID int64) ([]model.User, error)
}

type pgxChatRepository struct {
	q *queries.Queries
}

func NewChatRepository(q *queries.Queries) ChatRepository {
	return &pgxChatRepository{q: q}
}

func (r *pgxChatRepository) GetAll(ctx context.Context) ([]model.Chat, error) {
	c, err := r.q.ListChats(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]model.Chat, len(c))
	for i := range c {
		result[i] = dbChatToModel(c[i])
	}
	return result, nil
}

func (r *pgxChatRepository) GetForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error) {
	c, err := r.q.GetChatsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]model.Chat, len(c))
	for i := range c {
		u, err := r.q.GetChatParticipants(ctx, c[i].ID)
		if err != nil {
			return nil, err
		}

		chat := dbChatToModel(c[i])
		chat.Participants = dbUsersToModels(u)
		result[i] = chat
	}
	return result, nil
}

func (r *pgxChatRepository) GetByID(ctx context.Context, chatID int64) (*model.Chat, error) {
	c, err := r.q.GetChatById(ctx, chatID)
	if err != nil {
		return nil, err
	}

	u, err := r.q.GetChatParticipants(ctx, c.ID)
	if err != nil {
		return nil, err
	}

	chat := dbChatToModel(c)
	chat.Participants = dbUsersToModels(u)

	return &chat, nil
}

func (r *pgxChatRepository) GetMessagesForChat(ctx context.Context, chatID int64, limit int32, beforeTime time.Time) ([]model.Message, error) {
	m, err := r.q.GetMessagesForChat(ctx, queries.GetMessagesForChatParams{
		ChatID:    pgtype.Int8{Int64: chatID, Valid: true},
		Limit:     limit,
		CreatedAt: pgtype.Timestamptz{Time: beforeTime, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return dbMessagesToModels(m), nil
}

func (r *pgxChatRepository) Create(ctx context.Context) (model.Chat, error) {
	chat, err := r.q.CreateChat(ctx)
	if err != nil {
		return model.Chat{}, err
	}

	return dbChatToModel(chat), nil
}

func (r *pgxChatRepository) AddParticipant(ctx context.Context, chatID int64, userID uuid.UUID) error {
	return r.q.AddChatParticipant(ctx, queries.AddChatParticipantParams{
		ChatID: chatID,
		UserID: userID,
	})
}

func (r *pgxChatRepository) GetParticipants(ctx context.Context, chatID int64) ([]model.User, error) {
	u, err := r.q.GetChatParticipants(ctx, chatID)
	if err != nil {
		return nil, err
	}
	return dbUsersToModels(u), nil
}

func dbChatToModel(c queries.Chat) model.Chat {
	return model.Chat{
		ID:        c.ID,
		Title:     c.Title.String,
		CreatedAt: c.CreatedAt.Time,
		UpdatedAt: c.UpdatedAt.Time,
	}
}

func dbUserToModel(u queries.User) model.User {
	return model.User{
		ID:             u.ID,
		Name:           u.Name.String,
		Email:          u.Email,
		HashedPassword: u.HashedPassword.String,
		AvatarUrl:      u.AvatarUrl.String,
		Bio:            u.Bio.String,
		CreatedAt:      u.CreatedAt.Time,
		UpdatedAt:      u.UpdatedAt.Time,
	}
}

func dbUsersToModels(users []queries.User) []model.User {
	out := make([]model.User, len(users))
	for i := range users {
		out[i] = dbUserToModel(users[i])
	}
	return out
}

func dbMessagesToModels(msgs []queries.Message) []model.Message {
	out := make([]model.Message, len(msgs))
	for i := range msgs {
		senderID := func() *uuid.UUID {
			if !msgs[i].SenderID.Valid {
				return nil
			}
			uid, err := uuid.FromBytes(msgs[i].SenderID.Bytes[:])
			if err != nil {
				return nil
			}
			return &uid
		}()

		out[i] = model.Message{
			ID:       msgs[i].ID,
			ChatID:   msgs[i].ChatID.Int64,
			SenderID: senderID,
			Content:  msgs[i].Content.String,
			SentAt:   msgs[i].CreatedAt.Time,
			DeliveredAt: func() *time.Time {
				if msgs[i].DeliveredAt.Valid {
					return &msgs[i].DeliveredAt.Time
				}
				return nil
			}(),
			ReadAt: func() *time.Time {
				if msgs[i].ReadAt.Valid {
					return &msgs[i].ReadAt.Time
				}
				return nil
			}(),
		}
	}
	return out
}

var _ ChatRepository = (*pgxChatRepository)(nil)
