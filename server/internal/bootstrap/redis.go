package bootstrap

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"pycrs.cz/what-it-doo/internal/apiserver"
)

func InitRedis(config *apiserver.Configuration) (*redis.Client, error) {
	fmt.Println(config)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
