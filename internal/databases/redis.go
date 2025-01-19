package databases

import (
	"context"
	"fmt"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, cfg *configs.Config) (*redis.Client, error) {
	hostPort := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     hostPort,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
