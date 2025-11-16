package worker

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Task struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type MessagePayload struct {
	Content      string    `json:"content"`
	SenderID     uuid.UUID `json:"sender_id"`
	ChatID       int64     `json:"chat_id"`
	TempID       int64     `json:"temp_id"`
	ConnectionID uuid.UUID `json:"connection_id"`
}
