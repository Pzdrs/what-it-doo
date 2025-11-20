package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/bus/payload"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/ws"
)

func handleUserTyping(ctx context.Context, ev payload.Event, connectionManager ws.ConnectionManager, chatService service.ChatService) {
	var payload payload.UserTypingEventPayload
	if err := json.Unmarshal(ev.Payload, &payload); err != nil {
		log.Printf("Failed to unmarshal UserTypingPayload: %v", err)
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

		connectionManager.BroadcastUser(userID, map[string]any{
			"type": ws.UserTypingMessageType,
			"data": ws.TypingFanoutPayload{
				ChatID: payload.ChatID,
				UserID: payload.UserID,
				Typing: payload.Typing,
			},
		}, []uuid.UUID{payload.OriginConnectionID})
	}
}

func handleDapUp(ctx context.Context, connectionManager ws.ConnectionManager, chatService service.ChatService) {
	// TOOD: implement
}

func handlePresenceChange(ev payload.Event, connectionManager ws.ConnectionManager) {
	var payload payload.PresenceChangeEventPayload
	if err := json.Unmarshal(ev.Payload, &payload); err != nil {
		log.Printf("Failed to unmarshal PresenceChangeEventPayload: %v", err)
		return
	}

	connectionManager.Broadcast(map[string]any{
		"type": ws.PresenceChangeMessageType,
		"data": ws.PresenceChangeFanoutPayload{
			UserID: payload.UserID,
			Online: payload.Online,
		},
	}, []uuid.UUID{})
}
