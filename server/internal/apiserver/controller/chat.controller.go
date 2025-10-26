package controller

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal/apiserver/common"
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
	common.WriteJSON(w, 200, chats)
}
