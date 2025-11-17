package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
	"pycrs.cz/what-it-doo/internal/bus/payload"
)

func handleMessageAck(ctx context.Context, ev payload.Event, connectionManager ws.ConnectionManager, chatService service.ChatService) {
	var payload payload.MessageAckEventPayload
	if err := json.Unmarshal(ev.Payload, &payload); err != nil {
		log.Printf("Failed to unmarshal MessageAckPayload: %v", err)
		return
	}

	for _, userID := range connectionManager.GetConnectedUsers() {
		is, err := chatService.IsUserInChat(ctx, userID, payload.ChatID)
		if err != nil {
			log.Printf("Failed to check if user %s is in chat %d: %v", userID, payload.ChatID, err)
			return
		}
		if !is {
			continue
		}

		conns := connectionManager.GetUserConnections(userID)
		for _, conn := range conns {
			if conn.ID == payload.ConnectionID {
				msg := map[string]interface{}{
					"type": ws.MessageAckMessageType,
					"data": ws.ChatMessageAck{
						TempID:    payload.TempID,
						MessageID: payload.MessageID,
						SentAt:    payload.SentAt,
					},
				}
				err := conn.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("Failed to send MessageAck to user %s: %v", userID, err)
				}
			}
		}
	}
}

// This could be optimized the fuck out of so we dont waste any DB lookups but I can't be fucked
func handleMessageFanout(ctx context.Context, ev payload.Event, connectionManager ws.ConnectionManager, chatService service.ChatService) {
	var payload payload.MessageFanoutEventPayload
	if err := json.Unmarshal(ev.Payload, &payload); err != nil {
		log.Printf("Failed to unmarshal MessageFanoutPayload: %v", err)
		return
	}

	message, err := chatService.GetMessageByID(ctx, payload.MessageID)
	if err != nil {
		log.Printf("Failed to get message %d: %v", payload.MessageID, err)
		return
	}

	for _, userID := range connectionManager.GetConnectedUsers() {
		is, err := chatService.IsUserInChat(ctx, userID, payload.ChatID)
		if err != nil {
			log.Printf("Failed to check if user %s is in chat %d: %v", userID, payload.ChatID, err)
			return
		}
		if !is {
			continue
		}

		connectionManager.BroadcastToUser(userID, map[string]interface{}{
			"type": ws.NewMessageMessageType,
			"data": map[string]interface{}{
				"chat_id": payload.ChatID,
				"message": message,
			},
		}, []uuid.UUID{payload.OriginConnectionID})
	}
}
