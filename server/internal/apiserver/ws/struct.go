package ws

import (
	"encoding/json"
	"time"
)

const (
	TypeChatMessage    = "message"
	TypeChatMessageAck = "message_ack"
	TypeTyping         = "typing"
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

// ChatMessageAck represents an acknowledgment of a received chat message
type ChatMessageAck struct {
	TempID int64     `json:"temp_id"`
	SentAt time.Time `json:"sent_at"`
}

type TypingMessage struct {
	Typing bool `json:"typing"`
}

func NewMessage(typ string, payload interface{}) BaseMessage {
	json, err := json.Marshal(payload)
	if err != nil {
		panic("failed to marshal message")
	}
	return BaseMessage{
		Type: typ,
		Data: json,
	}
}
