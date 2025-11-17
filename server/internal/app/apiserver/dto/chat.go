package dto

import (
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/domain/model"
)

type Chat struct {
	ID           int64         `json:"id"`
	Title        string        `json:"title"`
	Participants []UserDetails `json:"participants"`
	LastMessage  string        `json:"last_message,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type ChatMessage struct {
	ID          int64      `json:"id"`
	SenderID    *uuid.UUID `json:"sender_id"`
	Content     string     `json:"content"`
	SentAt      time.Time  `json:"sent_at"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
}

type ChatMessages struct {
	Messages []ChatMessage `json:"messages"`
	HasMore  bool          `json:"has_more"`
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
