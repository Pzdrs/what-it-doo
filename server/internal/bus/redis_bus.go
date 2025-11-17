package bus

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"pycrs.cz/what-it-doo/internal/bus/payload"
)

const (
	taskStream        = "stream:tasks"
	taskConsumerGroup = "workers"
)

type redisCommunicationBus struct {
	r *redis.Client
}

func NewRedisCommunicationBus(r *redis.Client) CommnunicationBus {
	return &redisCommunicationBus{r: r}
}

func (b *redisCommunicationBus) EnqueueTask(ctx context.Context, typ string, payload any) (string, error) {
	marshaledPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return b.r.XAdd(ctx, &redis.XAddArgs{
		Stream: taskStream,
		Values: map[string]interface{}{
			"type":    typ,
			"payload": marshaledPayload,
		},
	}).Result()
}

func (b *redisCommunicationBus) ConsumeTasks(ctx context.Context) (<-chan payload.Task, error) {
	out := make(chan payload.Task)

	go func() {
		defer close(out)

		hostname, _ := os.Hostname()

		for {
			if ctx.Err() != nil {
				return
			}

			resp, err := b.r.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    taskConsumerGroup,
				Consumer: hostname,
				Streams:  []string{taskStream, ">"},
			}).Result()

			if err != nil {
				if err == redis.Nil {
					continue
				}
				log.Printf("Worker redis read error: %v", err)
				continue
			}

			for _, msg := range resp[0].Messages {
				t, ok := msg.Values["type"].(string)
				if !ok {
					continue
				}

				raw := msg.Values["payload"].(string)

				out <- payload.Task{
					ID:      msg.ID,
					Type:    t,
					Payload: json.RawMessage(raw),
				}
			}
		}
	}()

	return out, nil
}

func (b *redisCommunicationBus) AckTask(ctx context.Context, taskId any) {
	b.r.XAck(ctx, taskStream, taskConsumerGroup, taskId.(string))
	b.r.XDel(ctx, taskStream, taskId.(string))
}

func (b *redisCommunicationBus) DispatchGatewayEvent(ctx context.Context, gatewayId string, typ string, p any) error {
	log.Printf("Dispatching event to gateway %s: %s %+v", gatewayId, typ, p)
	event := map[string]interface{}{
		"type":    typ,
		"payload": p,
	}

	marshaledPayload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return b.r.Publish(ctx, "events:"+gatewayId, marshaledPayload).Err()
}

func (b *redisCommunicationBus) SubscribeGatewayEvents(ctx context.Context, gatewayId string) (<-chan payload.Event, error) {
	log.Printf("Subscribing to gateway events for %s", gatewayId)
	sub := b.r.Subscribe(ctx, "events:"+gatewayId)
	ch := sub.Channel()
	out := make(chan payload.Event)

	go func() {
		defer close(out)
		for msg := range ch {
			var event payload.Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("bad payload: %v", err)
				continue
			}

			out <- event
		}
	}()

	return out, nil
}

func (b *redisCommunicationBus) DispatchGlobalEvent(ctx context.Context, typ string, p any) error {
	log.Printf("Dispatching global event: %s %+v", typ, p)
	event := map[string]interface{}{
		"type":    typ,
		"payload": p,
	}

	marshaledPayload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return b.r.Publish(ctx, "events:global", marshaledPayload).Err()
}

func (b *redisCommunicationBus) SubscribeGlobalEvents(ctx context.Context) (<-chan payload.Event, error) {
	log.Print("Subscribing to global events")
	sub := b.r.Subscribe(ctx, "events:global")
	ch := sub.Channel()
	out := make(chan payload.Event)

	go func() {
		defer close(out)
		for msg := range ch {
			var event payload.Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("bad payload: %v", err)
				continue
			}

			out <- event
		}
	}()

	return out, nil
}
