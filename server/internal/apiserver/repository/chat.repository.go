package repository

import (
	"context"

	"pycrs.cz/what-it-do/internal/database"
)

type ChatRepository struct {
	q *database.Queries
}

func NewChatRepository(q *database.Queries) *ChatRepository {
	return &ChatRepository{q: q}
}

func (r *ChatRepository) GetAllChats() ([]database.Chat, error) {
	return r.q.ListChats(context.Background())
}
