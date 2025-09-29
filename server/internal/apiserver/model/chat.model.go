package model

import "time"

type Chat struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Chat) IsGroupChat() bool {
	return false
}