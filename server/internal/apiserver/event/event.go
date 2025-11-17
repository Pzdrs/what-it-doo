package event

import (
	"context"
	"log"

	"pycrs.cz/what-it-doo/internal/apiserver/service"
	"pycrs.cz/what-it-doo/internal/apiserver/ws"
	"pycrs.cz/what-it-doo/internal/bus"
)

func StartGatewayEventHandler(ctx context.Context, bus bus.CommnunicationBus, gatewayID string, connectionManager ws.ConnectionManager, chatService service.ChatService) error {
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
				case "message_ack":
					handleMessageAck(ctx, ev, connectionManager, chatService)
				default:
					log.Printf("Unknown event type: %s", ev.Type)
				}
			}
		}
	}()

	return nil
}

func StartGlobalEventHandler(ctx context.Context, bus bus.CommnunicationBus, connectionManager ws.ConnectionManager, chatService service.ChatService) error {
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
				case "message_fanout":
					handleMessageFanout(ctx, ev, connectionManager, chatService)
				default:
					log.Printf("Unknown event type: %s", ev.Type)
				}
			}
		}
	}()

	return nil
}
