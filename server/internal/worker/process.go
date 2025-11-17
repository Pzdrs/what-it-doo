package worker

import (
	"context"
	"log"

	"pycrs.cz/what-it-doo/internal/apiserver/service"
	b "pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/bus/payload"
)

func ProcessMessageTask(ctx context.Context, chatService service.ChatService, bus b.CommnunicationBus, p payload.MessageTaskPayload) error {
	log.Printf("Processing message from user %s in chat %d: %s\n", p.SenderID, p.ChatID, p.Content)
	message, err := chatService.SendMessage(ctx, p.ChatID, p.SenderID, p.Content)
	if err != nil {
		return err
	}

	if err := bus.DispatchGatewayEvent(ctx, p.GatewayID, b.MessageAckEventType, payload.MessageAckEventPayload{
		ConnectionID: p.ConnectionID,
		ChatID:       p.ChatID,
		TempID:       p.TempID,
		MessageID:    message.ID,
		SentAt:       message.SentAt,
	}); err != nil {
		return err
	}

	if err := bus.DispatchGlobalEvent(ctx, b.MessageFanoutEventType, payload.MessageFanoutEventPayload{
		ChatID:             p.ChatID,
		MessageID:          message.ID,
		OriginConnectionID: p.ConnectionID,
	}); err != nil {
		return err
	}

	return nil
}
