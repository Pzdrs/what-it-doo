// Payload definitions for WebSocket communication between clients and gateways.
package ws

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	NewMessageMessageType = "new_message"
	MessageAckMessageType = "message_ack"
	UserTypingMessageType = "typing"
	DapUpMessageType      = "dap_up"
)

type BaseMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// ChatMessage represents a chat message sent by the client
type ChatMessage struct {
	Message string `json:"message"`
	ChatID  int64  `json:"chat_id"`
	TempID  int64  `json:"temp_id"`
}

// ChatMessageAck represents an acknowledgment of a received chat message from the server
type ChatMessageAck struct {
	TempID    int64     `json:"temp_id"`
	MessageID int64     `json:"message_id"`
	SentAt    time.Time `json:"sent_at"`
}

type TypingPayload struct {
	Typing bool  `json:"typing"`
	ChatID int64 `json:"chat_id"`
}

type TypingFanoutPayload struct {
	ChatID int64     `json:"chat_id"`
	UserID uuid.UUID `json:"user_id"`
	Typing bool      `json:"typing"`
}
