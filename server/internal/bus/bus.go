package bus

import (
	"context"

	"pycrs.cz/what-it-doo/internal/bus/payload"
)

const (
	MessageTaskType = "message"
)

const (
	MessageAckEventType    = "message_ack"
	MessageFanoutEventType = "message_fanout"
	UserTypingEventType    = "typing"
)

type CommnunicationBus interface {
	// EnqueueTask dispatches a task to be processed by a worker. It returns the ID of the dispatched task.
	EnqueueTask(ctx context.Context, typ string, payload any) (string, error)
	// ConsumeTasks returns a channel that yields tasks to be processed by the worker.
	ConsumeTasks(ctx context.Context) (<-chan payload.Task, error)
	// AckTask acknowledges the completion of a task with the given ID.
	AckTask(ctx context.Context, taskId any)

	DispatchGlobalEvent(ctx context.Context, typ string, payload any) error
	SubscribeGlobalEvents(ctx context.Context) (<-chan payload.Event, error)

	// DispatchGatewayEvent dispatches an event to a given gateway.
	DispatchGatewayEvent(ctx context.Context, gatewayId string, typ string, payload any) error
	SubscribeGatewayEvents(ctx context.Context, gatewayId string) (<-chan payload.Event, error)
}
