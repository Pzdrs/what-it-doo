package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/common"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/config"
)

type ChatService interface {
	// GetChatsForUser retrieves all chats for a specific user.
	GetChatsForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error)
	// GetChatByID retrieves a chat by its ID.
	GetChatByID(ctx context.Context, chatID int64) (*model.Chat, error)
	// GetChatMessages retrieves messages for a specific chat.
	// It returns a slice of ChatMessage, a boolean indicating if there are more messages, and an error if any
	GetMessagesForChat(ctx context.Context, chatID int64, limit int32, before time.Time) ([]model.Message, bool, error)
	// CreateChat creates a new chat with the given participants.
	CreateChat(ctx context.Context, participants []string) (*model.Chat, error)
	// SendMessage sends a message in a chat.
	SendMessage(ctx context.Context, chatID int64, senderID uuid.UUID, content string) (model.Message, error)
}

type chatService struct {
	repository     repository.ChatRepository
	userRepository repository.UserRepository
	config         config.Configuration
}

func NewChatService(repo repository.ChatRepository, userRepo repository.UserRepository, config config.Configuration) ChatService {
	return &chatService{
		repository:     repo,
		userRepository: userRepo,
		config:         config,
	}
}

func (c *chatService) CreateChat(ctx context.Context, participants []string) (*model.Chat, error) {
	chat, err := c.repository.Create(ctx)
	if err != nil {
		return nil, err
	}

	for _, participantEmail := range participants {
		u, err := c.userRepository.GetByEmail(ctx, participantEmail)
		if err != nil {
			return nil, err
		}
		if err = c.repository.AddParticipant(ctx, chat.ID, u.ID); err != nil {
			return nil, err
		}
	}

	chatParticipants, err := c.repository.GetParticipants(ctx, chat.ID)
	if err != nil {
		return nil, err
	}
	chat.Participants = chatParticipants

	return &chat, nil
}

func (s *chatService) GetChatsForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error) {
	chats, err := s.repository.GetForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i := range chats {
		for j := range chats[i].Participants {
			chats[i].Participants[j].AvatarUrl = common.GetAvatarUrl(chats[i].Participants[j], s.config.Gravatar)
		}
	}
	return chats, nil
}

func (s *chatService) GetChatByID(ctx context.Context, chatID int64) (*model.Chat, error) {
	chat, err := s.repository.GetByID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	for i := range chat.Participants {
		chat.Participants[i].AvatarUrl = common.GetAvatarUrl(chat.Participants[i], s.config.Gravatar)
	}

	return chat, nil
}

func (s *chatService) GetMessagesForChat(ctx context.Context, chatID int64, limit int32, before time.Time) ([]model.Message, bool, error) {
	messages, err := s.repository.GetMessagesForChat(ctx, chatID, limit+1, before)
	if err != nil {
		return nil, false, err
	}

	hasMore := len(messages) > int(limit)

	if hasMore {
		messages = messages[:limit]
	}

	return messages, hasMore, nil
}

func (s *chatService) SendMessage(ctx context.Context, chatID int64, senderID uuid.UUID, content string) (model.Message, error) {
	return s.repository.CreateMessage(ctx, chatID, senderID, content)
}

var _ ChatService = (*chatService)(nil)
