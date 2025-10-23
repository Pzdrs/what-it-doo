package controller

import (
	"encoding/json"
	"net/http"

	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

type ChatController struct {
	chatService *service.ChatService
}

func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{chatService: chatService}
}

func (c *ChatController) HandleAllChats(w http.ResponseWriter, r *http.Request) {
	chats, err := c.chatService.GetAllChats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chats)
}
