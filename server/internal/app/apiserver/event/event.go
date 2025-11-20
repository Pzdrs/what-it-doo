package event

import (
	"context"
	"log"

	b "pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/ws"
)

func StartGatewayEventHandler(ctx context.Context, bus b.CommnunicationBus, gatewayID string, connectionManager ws.ConnectionManager, chatService service.ChatService) error {
	eventsCh, err := bus.SubscribeGatewayEvents(ctx, gatewayID)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Gateway event handler for %s shutting down", gatewayID)
				return
			case ev, ok := <-eventsCh:
				if !ok {
					log.Printf("Gateway events channel closed for %s", gatewayID)
					return
				}

				switch ev.Type {
				case b.MessageAckEventType:
					handleMessageAck(ctx, ev, connectionManager, chatService)
				default:
					log.Printf("Received unknown event type: %s", ev.Type)
				}
			}
		}
	}()

	return nil
}

func StartGlobalEventHandler(ctx context.Context, bus b.CommnunicationBus, connectionManager ws.ConnectionManager, chatService service.ChatService) error {
	eventsCh, err := bus.SubscribeGlobalEvents(ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Print("Global event handler shutting down")
				return
			case ev, ok := <-eventsCh:
				if !ok {
					log.Print("Global events channel closed")
					return
				}

				switch ev.Type {
				case b.MessageFanoutEventType:
					handleMessageFanout(ctx, ev, connectionManager, chatService)
				case b.UserTypingEventType:
					handleUserTyping(ctx, ev, connectionManager, chatService)
				case b.DapUpEventType:
					handleDapUp(ctx, connectionManager, chatService)
				case b.PresenceChangeEventType:
					handlePresenceChange(ev, connectionManager)
				default:
					log.Printf("Received unknown event type: %s", ev.Type)
				}
			}
		}
	}()

	return nil
}
