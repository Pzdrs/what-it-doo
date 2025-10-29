package model

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Message struct {
	ID          uuid.UUID  `json:"id"`
	SenderID    *uuid.UUID `json:"sender_id"`
	Content     string     `json:"content"`
	SentAt      time.Time  `json:"sent_at"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
}

func (c *Chat) IsGroupChat() bool {
	return false
}
