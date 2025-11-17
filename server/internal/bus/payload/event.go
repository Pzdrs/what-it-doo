package payload

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	MessageAckEvent = "message_ack"
	MessageFanoutEvent = "message_fanout"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type MessageAckEventPayload struct {
	ConnectionID uuid.UUID `json:"connection_id"`
	ChatID       int64     `json:"chat_id"`
	TempID       int64     `json:"temp_id"`
	MessageID    int64     `json:"message_id"`
	SentAt       time.Time `json:"sent_at"`
}

type MessageFanoutEventPayload struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}