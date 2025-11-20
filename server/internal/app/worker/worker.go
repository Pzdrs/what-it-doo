package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"pycrs.cz/what-it-doo/internal/app/worker/processor"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/bus/payload"
	"pycrs.cz/what-it-doo/internal/config"
	"pycrs.cz/what-it-doo/internal/domain/repository"
	"pycrs.cz/what-it-doo/internal/domain/service"
	"pycrs.cz/what-it-doo/internal/queries"
)

func RunWorker(ctx context.Context, q *queries.Queries, bus bus.CommnunicationBus, config config.Configuration) error {
	userRepository := repository.NewUserRepository(q)
	chatRepository := repository.NewChatRepository(q)

	chatService := service.NewChatService(chatRepository, userRepository, config)

	taskChan, err := bus.ConsumeTasks(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case task := <-taskChan:
			switch task.Type {
			case "message":
				var payload payload.MessageTaskPayload
				if err := json.Unmarshal(task.Payload, &payload); err != nil {
					fmt.Println("❌ Failed to unmarshal message payload:", err)
					continue
				}

				if err := processor.ProcessMessageTask(ctx, chatService, bus, payload); err != nil {
					fmt.Println("❌ Failed to process message task:", err)
					continue
				}
				fmt.Println("✅ Message task processed successfully")
			default:
				fmt.Println("⚠️ Unknown task type:", task.Type)
			}

			bus.AckTask(ctx, task.ID)
		}
	}
}
