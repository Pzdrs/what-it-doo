package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"pycrs.cz/what-it-doo/internal/apiserver/middleware"
	"pycrs.cz/what-it-doo/internal/apiserver/problem"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
	"pycrs.cz/what-it-doo/internal/worker"
)

type SocketController struct {
	ctx context.Context

	upgrader          websocket.Upgrader
	connectionManager ws.ConnectionManager
	redisClient       *redis.Client
	userService       service.UserService
}

func NewSocketController(ctx context.Context, upgrader websocket.Upgrader, connectionManager ws.ConnectionManager, redisClient *redis.Client, userService service.UserService) *SocketController {
	return &SocketController{ctx: ctx, upgrader: upgrader, connectionManager: connectionManager, redisClient: redisClient, userService: userService}
}

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

	connId := c.connectionManager.AddConnection(user.ID, session.ID, conn)

	go func() {
		defer conn.Close()
		defer c.connectionManager.RemoveConnection(user.ID, session.ID, connId)

		for {
			_, msgBytes, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("WebSocket closed:", err)
				break
			}

			fmt.Println("Raw message:", string(msgBytes))

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

				payload, err := json.Marshal(worker.MessagePayload{
					Content:      chatMessage.Message,
					SenderID:     user.ID,
					TempID:       chatMessage.TempID,
					ChatID:       chatMessage.ChatID,
					ConnectionID: connId,
				})
				if err != nil {
					fmt.Println("❌ Failed to marshal message payload:", err)
					continue
				}

				res, err := c.redisClient.XAdd(c.ctx, &redis.XAddArgs{
					Stream: "stream:tasks",
					Values: map[string]interface{}{
						"type":    "message",
						"payload": payload,
					},
				}).Result()

				if err != nil {
					fmt.Println("❌ Failed to enqueue message task:", err)
					continue
				}

				log.Println("Enqueued message task with ID:", res)
				// TODO: move sending ack to worker after message is stored in DB
				// conn.WriteJSON(ws.NewMessage(ws.TypeChatMessageAck, ws.ChatMessageAck{
				// 	TempID: chatMessage.TempID,
				// 	SentAt: timestamp,
				// }))

			default:
				log.Println("⚠️ Unknown message type:", base.Type)
			}
		}
	}()
}
