package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
)

type ChatService interface {
	// GetAllChats retrieves all chats.
	GetAllChats(ctx context.Context) ([]model.Chat, error)
	// GetChatsForUser retrieves all chats for a specific user.
	GetChatsForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error)
	// GetChatByID retrieves a chat by its ID.
	GetChatByID(ctx context.Context, chatID int64) (*model.Chat, error)
	// GetChatMessages retrieves messages for a specific chat.
	// It returns a slice of ChatMessage, a boolean indicating if there are more messages, and an error if any
	GetMessagesForChat(ctx context.Context, chatID int64, limit int32, before time.Time) ([]model.Message, bool, error)
	// CreateChat creates a new chat with the given participants.
	CreateChat(ctx context.Context, participants []string) (*model.Chat, error)
}

type chatService struct {
	repository     repository.ChatRepository
	userRepository repository.UserRepository
}

func NewChatService(repo repository.ChatRepository, userRepo repository.UserRepository) ChatService {
	return &chatService{
		repository:     repo,
		userRepository: userRepo,
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

func (s *chatService) GetAllChats(ctx context.Context) ([]model.Chat, error) {
	return s.repository.GetAll(ctx)
}

func (s *chatService) GetChatsForUser(ctx context.Context, userID uuid.UUID) ([]model.Chat, error) {
	return s.repository.GetForUser(ctx, userID)
}

func (s *chatService) GetChatByID(ctx context.Context, chatID int64) (*model.Chat, error) {
	return s.repository.GetByID(ctx, chatID)
}

func (s *chatService) GetMessagesForChat(ctx context.Context, chatID int64, limit int32, before time.Time) ([]model.Message, bool, error) {
	messages, err := s.repository.GetMessagesForChat(ctx, chatID, limit+1, before)
	if err != nil {
		return nil, false, err
	}
	return messages, len(messages) > int(limit), nil
}

var _ ChatService = (*chatService)(nil)
