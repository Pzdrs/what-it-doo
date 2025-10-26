package repository

import (
	"context"

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
