package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/dto"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/queries"
)

type ChatService interface {
	// GetAllChats retrieves all chats.
	GetAllChats() ([]model.Chat, error)
	// GetChatsForUser retrieves all chats for a specific user.
	GetChatsForUser(userID uuid.UUID) ([]dto.Chat, error)
	// GetChatByID retrieves a chat by its ID.
	GetChatByID(chatID int64) (*dto.Chat, error)
	// GetChatMessages retrieves messages for a specific chat.
	// It returns a slice of ChatMessage, a boolean indicating if there are more messages, and an error if any
	GetMessagesForChat(chatID int64, limit int32, before time.Time) ([]model.Message, bool, error)
}

type chatService struct {
	repository *repository.ChatRepository
}

func NewChatService(repo *repository.ChatRepository) ChatService {
	return &chatService{
		repository: repo,
	}
}

func mapChatToModel(chat queries.Chat) model.Chat {
	return model.Chat{
		ID:        chat.ID,
		Title:     chat.Title.String,
		CreatedAt: chat.CreatedAt.Time,
		UpdatedAt: chat.UpdatedAt.Time,
	}
}

func mapMessageToModel(msg queries.Message) model.Message {
	return model.Message{
		ID: msg.ID,
		SenderID: func() *uuid.UUID {
			if !msg.SenderID.Valid {
				return nil
			}
			u, err := uuid.FromBytes(msg.SenderID.Bytes[:])
			if err != nil {
				return nil
			}
			return &u
		}(),
		Content: msg.Content.String,
		SentAt:  msg.CreatedAt.Time,
		DeliveredAt: func() *time.Time {
			if msg.DeliveredAt.Valid {
				return &msg.DeliveredAt.Time
			} else {
				return nil
			}
		}(),
		ReadAt: func() *time.Time {
			if msg.ReadAt.Valid {
				return &msg.ReadAt.Time
			} else {
				return nil
			}
		}(),
	}
}

func (s *chatService) GetAllChats() ([]model.Chat, error) {
	return func() ([]model.Chat, error) {
		chats, err := s.repository.GetAllChats()
		if err != nil {
			return nil, err
		}
		result := []model.Chat{}
		for _, chat := range chats {
			result = append(result, mapChatToModel(chat))
		}
		return result, nil
	}()
}

func (s *chatService) GetChatsForUser(userID uuid.UUID) ([]dto.Chat, error) {
	return func() ([]dto.Chat, error) {
		chats, err := s.repository.GetChatsForUserWithParticipants(userID)
		if err != nil {
			return nil, err
		}
		result := []dto.Chat{}
		for _, chat := range chats {
			var participants []dto.UserDetails
			if len(chat.Participants) > 0 {
				if err := json.Unmarshal(chat.Participants, &participants); err != nil {
					return nil, fmt.Errorf("failed to unmarshal participants for chat %s: %w", chat.ID, err)
				}
			}

			result = append(result, dto.Chat{
				ID:           chat.ID,
				Title:        chat.Title.String,
				CreatedAt:    chat.CreatedAt.Time,
				UpdatedAt:    chat.UpdatedAt.Time,
				Participants: participants,
			})
		}
		return result, nil
	}()
}

func (s *chatService) GetChatByID(chatID int64) (*dto.Chat, error) {
	row, err := s.repository.GetChatByID(chatID)
	if err != nil {
		return nil, err
	}
	chat := dto.Chat{
		ID:        row.ID,
		Title:     row.Title.String,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
	json.Unmarshal(row.Participants, &chat.Participants)
	return &chat, nil
}

func (s *chatService) GetMessagesForChat(chatID int64, limit int32, before time.Time) ([]model.Message, bool, error) {
	messages, err := s.repository.GetMessagesForChat(chatID, limit+1, before)
	if err != nil {
		return nil, false, err
	}
	result := []model.Message{}
	for _, msg := range messages {
		result = append(result, mapMessageToModel(msg))
	}
	return result, len(messages) > int(limit), nil
}
