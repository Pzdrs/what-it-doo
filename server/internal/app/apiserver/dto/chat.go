package dto

import (
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/domain/model"
)

type Chat struct {
	ID           int64         `json:"id" validate:"required"`
	Title        string        `json:"title"`
	Participants []UserDetails `json:"participants" validate:"required"`
	LastMessage  string        `json:"last_message"`
	CreatedAt    time.Time     `json:"created_at" validate:"required"`
	UpdatedAt    time.Time     `json:"updated_at" validate:"required"`
}

type ChatMessage struct {
	ID          int64      `json:"id" validate:"required"`
	SenderID    *uuid.UUID `json:"sender_id"`
	Content     string     `json:"content" validate:"required"`
	SentAt      time.Time  `json:"sent_at"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
}

type ChatMessages struct {
	Messages []ChatMessage `json:"messages" validate:"required"`
	HasMore  bool          `json:"has_more" validate:"required"`
}

type CreateChatRequest struct {
	Participants []string `json:"participants" validate:"required,min=1"`
}

func MapMessageToDTO(message model.Message) ChatMessage {
	return ChatMessage{
		ID:          message.ID,
		SenderID:    message.SenderID,
		Content:     message.Content,
		SentAt:      message.SentAt,
		DeliveredAt: message.DeliveredAt,
		ReadAt:      message.ReadAt,
	}
}
