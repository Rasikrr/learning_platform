package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
	SetWithExpiration(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	val, err := r.redisStringCmd(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value any) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) SetWithExpiration(ctx context.Context, key string, value any, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) redisStringCmd(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}
