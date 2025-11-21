package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"pycrs.cz/what-it-doo/internal/app/apiserver/common"
	"pycrs.cz/what-it-doo/internal/app/apiserver/dto"
	"pycrs.cz/what-it-doo/internal/app/apiserver/http/middleware"
	"pycrs.cz/what-it-doo/internal/app/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/service"
)

type ChatController struct {
	chatService service.ChatService
	config      config.Configuration
}

func NewChatController(chatService service.ChatService, config config.Configuration) *ChatController {
	return &ChatController{chatService: chatService, config: config}
}

// HandleMyChats retrieves all chats the authenticated user is a participant of
//
//	@Summary		Get my chats
//	@Id				GetMyChats
//	@Description	Retrieves all chats the authenticated user is a participant of
//	@Tags			chats
//	@Produce		json
//	@Success		200	{array}	dto.Chat
//	@Security		SessionAuth
//	@Router			/chats/ [get]
func (c *ChatController) HandleMyChats(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionFromContext(r.Context())

	chats, err := c.chatService.GetChatsForUser(r.Context(), session.UserID)
	if err != nil {
		problem.Write(w, problem.NewInternalServerError(r, err))
		return
	}
	common.Encode(w, r, 200, chats)
}

// HandleGetChat retrieves a chat by its ID
//
//	@Summary		Get chat by ID
//	@Id				GetChatByID
//	@Description	Retrieves a chat by its ID
//	@Tags			chats
//	@Produce		json
//	@Param			chat_id	path		int	true	"Chat ID"
//	@Success		200		{object}	dto.Chat
//	@Failure		404
//	@Router			/chats/{chat_id} [get]
func (c *ChatController) HandleGetChat(w http.ResponseWriter, r *http.Request) {
	chat_id, err := strconv.ParseInt(chi.URLParam(r, "chat_id"), 10, 64)
	if err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid chat ID",
			"The provided chat ID is not a valid integer",
			"chats/invalid-chat-id",
		))
		return
	}

	chat, err := c.chatService.GetChatByID(r.Context(), chat_id)
	if errors.Is(err, sql.ErrNoRows) {
		problem.Write(w, problem.New(
			r, http.StatusNotFound,
			"Chat not found",
			"No chat found with the specified ID",
			"chats/chat-not-found",
		))
		return
	} else if err != nil {
		problem.Write(w, problem.NewInternalServerError(r, err))
		return
	}
	common.Encode(w, r, 200, chat)
}

// HandleCreateChat creates a new chat
//
//	@Summary		Create a new chat
//	@Id				CreateChat
//	@Description	Creates a new chat with the specified participants
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateChatRequest	true	"Chat creation request"
//	@Success		201		{object}	dto.Chat
//	@Router			/chats/ [post]
func (c *ChatController) HandleCreateChat(w http.ResponseWriter, r *http.Request) {
	req, ok := common.DecodeValidate[dto.CreateChatRequest](w, r)
	if !ok {
		return
	}

	chat, err := c.chatService.CreateChat(r.Context(), req.Participants)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			problem.Write(w, problem.New(
				r, http.StatusBadRequest,
				"Invalid participant",
				"One or more specified participants do not exist",
				"chats/invalid-participant",
			))
		} else {
			problem.Write(w, problem.NewInternalServerError(r, err))
		}
		return
	}
	common.Encode(w, r, 201, chat)
}

// HandleGetChatMessages retrieves messages for a specific chat
//
//	@Summary		Get chat messages
//	@Id				GetChatMessages
//	@Description	Retrieves messages for a specific chat
//	@Tags			chats
//	@Produce		json
//	@Param			chat_id	path		int		true	"Chat ID"
//	@Param			limit	query		int		false	"Maximum number of messages to retrieve"		default(50)
//	@Param			before	query		string	false	"Retrieve messages sent before this timestamp"	format(date-time)	example(2023-01-01T00:00:00Z)
//	@Success		200		{object}	dto.ChatMessages
//	@Router			/chats/{chat_id}/messages [get]
func (c *ChatController) HandleGetChatMessages(w http.ResponseWriter, r *http.Request) {
	chat_id, err := strconv.ParseInt(chi.URLParam(r, "chat_id"), 10, 64)
	if err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid chat ID",
			"The provided chat ID is not a valid integer",
			"chats/invalid-chat-id",
		))
		return
	}

	limit, err := common.ParseQueryInt[int32](r, "limit", 50)
	if err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid limit",
			"The provided limit is not a valid integer",
			"chats/invalid-limit",
		))
		return
	}

	before, err := common.ParseQueryTime(r, "before", time.Now())
	if err != nil {
		problem.Write(w, problem.New(
			r, http.StatusBadRequest,
			"Invalid before",
			"The provided before timestamp is not valid",
			"chats/invalid-before",
		))
		return
	}

	messages, more, err := c.chatService.GetMessagesForChat(r.Context(), chat_id, limit, before)
	if err != nil {
		problem.Write(w, problem.NewInternalServerError(r, err))
		return
	}
	common.Encode(w, r, 200, dto.ChatMessages{
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
