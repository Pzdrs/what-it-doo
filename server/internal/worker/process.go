package worker

import (
	"context"
	"fmt"

	"pycrs.cz/what-it-doo/internal/apiserver/service"
)

func ProcessMessageTask(ctx context.Context, chatService service.ChatService, payload MessagePayload) error {
	fmt.Printf("Processing message from user %s in chat %d: %s\n", payload.SenderID, payload.ChatID, payload.Content)
	fmt.Println(payload)
	message, err := chatService.SendMessage(ctx, payload.ChatID, payload.SenderID, payload.Content)
	_ = message
	if err != nil {
		return err
	}

	// TODO: send ack to the sending connection

	return nil
}
