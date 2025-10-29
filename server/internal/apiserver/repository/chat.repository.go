package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"pycrs.cz/what-it-doo/internal/queries"
)

type ChatRepository struct {
	q *queries.Queries
}

func NewChatRepository(q *queries.Queries) *ChatRepository {
	return &ChatRepository{q: q}
}

func (r *ChatRepository) GetAllChats() ([]queries.Chat, error) {
	return r.q.ListChats(context.Background())
}

func (r *ChatRepository) GetChatsForUser(userID uuid.UUID) ([]queries.Chat, error) {
	return r.q.GetChatsForUser(context.Background(), userID)
}

func (r *ChatRepository) GetChatsForUserWithParticipants(userID uuid.UUID) ([]queries.GetChatsForUserWithParticipantsRow, error) {
	return r.q.GetChatsForUserWithParticipants(context.Background(), userID)
}

func (r *ChatRepository) GetChatByID(chatID int64) (*queries.GetChatByIdWithParticipantsRow, error) {
	chat, err := r.q.GetChatByIdWithParticipants(context.Background(), chatID)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) GetMessagesForChat(chatID int64, limit int32, beforeTime time.Time) ([]queries.Message, error) {
	return r.q.GetMessagesForChat(context.Background(), queries.GetMessagesForChatParams{
		ChatID:    pgtype.Int8{Int64: chatID, Valid: true},
		Limit:     limit,
		CreatedAt: pgtype.Timestamptz{Time: beforeTime, Valid: true},
	})
}
