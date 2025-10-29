package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"pycrs.cz/what-it-doo/internal/apiserver/common"
	"pycrs.cz/what-it-doo/internal/apiserver/dto"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

type ChatController struct {
	chatService service.ChatService
}

func NewChatController(chatService service.ChatService) *ChatController {
	return &ChatController{chatService: chatService}
}

// HandleMyChats
//
//	@Summary		Get my chats
//	@Id				GetMyChats
//	@Description	Retrieves all chats the authenticated user is a participant of
//	@Tags			Chats
//	@Produce		json
//	@Success		200	{array}	dto.Chat
//	@Security		SessionAuth
//	@Router			/chats/ [get]
func (c *ChatController) HandleMyChats(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionFromContext(r.Context())

	chats, err := c.chatService.GetChatsForUser(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, 200, chats)
}

// HandleGetChat
//
//	@Summary		Get chat by ID
//	@Id				GetChatByID
//	@Description	Retrieves a chat by its ID
//	@Tags			Chats
//	@Produce		json
//	@Param			chat_id	path		int	true	"Chat ID"
//	@Success		200		{object}	dto.Chat
//	@Failure		404
//	@Router			/chats/{chat_id} [get]
func (c *ChatController) HandleGetChat(w http.ResponseWriter, r *http.Request) {
	chat_id, err := strconv.ParseInt(chi.URLParam(r, "chat_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	chat, err := c.chatService.GetChatByID(chat_id)
	if errors.Is(err, sql.ErrNoRows) {
		common.WriteJSON(w, 404, struct{}{})
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, 200, chat)
}

// HandleGetChatMessages
//
//	@Summary		Get chat messages
//	@Id				GetChatMessages
//	@Description	Retrieves messages for a specific chat
//	@Tags			Chats
//	@Produce		json
//	@Param			chat_id	path		int		true	"Chat ID"
//	@Param			limit	query		int		false	"Maximum number of messages to retrieve"		default(50)
//	@Param			before	query		string	false	"Retrieve messages sent before this timestamp"	format(date-time)	example(2023-01-01T00:00:00Z)
//	@Success		200		{object}	dto.ChatMessages
//	@Router			/chats/{chat_id}/messages [get]
func (c *ChatController) HandleGetChatMessages(w http.ResponseWriter, r *http.Request) {
	chat_id, err := strconv.ParseInt(chi.URLParam(r, "chat_id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	limit, err := common.ParseQueryInt[int32](r, "limit", 50)
	if err != nil {
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}

	before, err := common.ParseQueryTime(r, "before", time.Now())
	if err != nil {
		http.Error(w, "Invalid before", http.StatusBadRequest)
		return
	}

	messages, more, err := c.chatService.GetMessagesForChat(chat_id, limit, before)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, 200, dto.ChatMessages{
		Messages: func() []dto.ChatMessage {
			result := []dto.ChatMessage{}
			for _, msg := range messages {
				result = append(result, dto.MapMessageToDTO(msg))
			}
			return result
		}(),
		HasMore: more,
	})
}
