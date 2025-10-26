package service

import (
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/queries"
)

type ChatService struct {
	repository *repository.ChatRepository
}

func NewChatService(repo *repository.ChatRepository) *ChatService {
	return &ChatService{
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

func (s *ChatService) GetAllChats() ([]model.Chat, error) {
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
