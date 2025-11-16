package bootstrap

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func EnsureStreamGroup(ctx context.Context, redis *redis.Client) {
	log.Println("Ensuring stream groups exist")
	redis.XGroupCreateMkStream(ctx, "stream:tasks", "workers", "$")
}
