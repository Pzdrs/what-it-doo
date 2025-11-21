package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"pycrs.cz/what-it-doo/internal/app/apiserver/http/middleware"
	"pycrs.cz/what-it-doo/internal/app/apiserver/presence"
	"pycrs.cz/what-it-doo/internal/app/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/bus"
	b "pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/bus/payload"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/ws"
)

type SocketController struct {
	ctx context.Context

	upgrader          websocket.Upgrader
	connectionManager ws.ConnectionManager
	presenceManager   *presence.PresenceManager
	bus               b.CommnunicationBus
	userService       service.UserService
	gatewayID         string
}

func NewSocketController(ctx context.Context, upgrader websocket.Upgrader, connectionManager ws.ConnectionManager, bus bus.CommnunicationBus, userService service.UserService, gatewayID string, presenceManager *presence.PresenceManager) *SocketController {
	return &SocketController{ctx: ctx, upgrader: upgrader, connectionManager: connectionManager, bus: bus, userService: userService, gatewayID: gatewayID, presenceManager: presenceManager}
}

// HandleWebSocket handles WebSocket connections
//
//	@Summary		Handle WebSocket connections
//	@Description	Establish a WebSocket connection for real-time communication
//	@Id				handleWebSocket
//	@Tags			miscellaneous
//	@Produce		json
//	@Success		101	{string}	string	"Switching Protocols"
//	@Failure		500	{object}	problem.Problem
//	@Router			/ws [get]
func (c *SocketController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.SessionFromContext(r.Context())
	user, err := c.userService.GetByID(r.Context(), session.UserID)

	if err != nil {
		problem.Write(w, problem.NewInternalServerError(r, err))
		return
	}

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	connID := c.connectionManager.AddConnection(user.ID, session.ID, conn)
	if err := c.presenceManager.AddConnection(c.ctx, user.ID); err != nil {
		problem.Write(w, problem.NewInternalServerError(r, err))
		return
	}

	go func() {
		defer conn.Close()
		defer c.connectionManager.RemoveConnection(user.ID, session.ID, connID)
		defer c.presenceManager.RemoveConnection(c.ctx, user.ID)

		for {
			_, msgBytes, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("WebSocket closed:", err)
				break
			}

			var base ws.BaseMessage
			if err := json.Unmarshal(msgBytes, &base); err != nil {
				fmt.Println("❌ Failed to parse base message:", err)
				continue
			}

			switch base.Type {

			case "message":
				var chatMessage ws.ChatMessage
				if err := json.Unmarshal(base.Data, &chatMessage); err != nil {
					fmt.Println("❌ Invalid chat message:", err)
					continue
				}

				taskId, err := c.bus.EnqueueTask(c.ctx, b.MessageTaskType, payload.MessageTaskPayload{
					Content:      chatMessage.Message,
					SenderID:     user.ID,
					TempID:       chatMessage.TempID,
					ChatID:       chatMessage.ChatID,
					ConnectionID: connID,
					GatewayID:    c.gatewayID,
				})

				if err != nil {
					fmt.Println("❌ Failed to enqueue message task:", err)
					continue
				}

				log.Println("Enqueued message task with ID:", taskId)
			case "typing":
				var typingPayload ws.TypingPayload
				if err := json.Unmarshal(base.Data, &typingPayload); err != nil {
					fmt.Println("❌ Invalid typing payload:", err)
					continue
				}

				if err := c.bus.DispatchGlobalEvent(c.ctx, b.UserTypingEventType, payload.UserTypingEventPayload{
					UserID:             user.ID,
					Typing:             typingPayload.Typing,
					ChatID:             typingPayload.ChatID,
					OriginConnectionID: connID,
				}); err != nil {
					fmt.Println("❌ Failed to dispatch typing event:", err)
					continue
				}
			case "dap_up":
				if err := c.bus.DispatchGlobalEvent(c.ctx, b.DapUpEventType, nil); err != nil {
					fmt.Println("❌ Failed to dispatch dap_up event:", err)
					continue
				}
			default:
				log.Println("⚠️ Unknown message type:", base.Type)
			}
		}
	}()
}
