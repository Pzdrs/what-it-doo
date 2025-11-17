package payload

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Task struct {
	ID      any             `json:"id"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type MessageTaskPayload struct {
	Content      string    `json:"content"`
	SenderID     uuid.UUID `json:"sender_id"`
	ChatID       int64     `json:"chat_id"`
	TempID       int64     `json:"temp_id"`
	ConnectionID uuid.UUID `json:"connection_id"`
	GatewayID    string    `json:"gateway_id"`
}
