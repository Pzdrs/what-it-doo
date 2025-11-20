// Payload definitions for events.
// Events are Worker -> Gateway or Gategway -> Gateway messages.
package payload

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// Sent by workers to the originating gateway to acknowledge a message has been processed and stored.
type MessageAckEventPayload struct {
	ConnectionID uuid.UUID `json:"connection_id"`
	ChatID       int64     `json:"chat_id"`
	TempID       int64     `json:"temp_id"`
	MessageID    int64     `json:"message_id"`
	SentAt       time.Time `json:"sent_at"`
}

// Sent by workers to all gateways to notify them of a new message in a chat.
type MessageFanoutEventPayload struct {
	ChatID             int64     `json:"chat_id"`
	MessageID          int64     `json:"message_id"`
	OriginConnectionID uuid.UUID `json:"origin_connection_id"`
}

// Sent by a gateway to all gateways to notify them that a user is typing.
type UserTypingEventPayload struct {
	Typing bool      `json:"typing"`
	ChatID int64     `json:"chat_id"`
	UserID uuid.UUID `json:"user_id"`
	// We need to track the origin to avoid echoing back to the user who is typing
	OriginConnectionID uuid.UUID `json:"origin_connection_id"`
}

type PresenceChangeEventPayload struct {
	UserID uuid.UUID `json:"user_id"`
	Online bool      `json:"online"`
}
