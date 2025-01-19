package auth

import (
	"context"
	"errors"
	cache2 "github.com/Rasikrr/learning_platform/internal/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	codeAndCredsLifeTime = time.Minute * 5
)

type Cache interface {
	GetCode(ctx context.Context, email string) (string, error)
	SetCode(ctx context.Context, email string, code string) error
	GetPasswordHash(ctx context.Context, email string) (string, error)
	SetPasswordHash(ctx context.Context, email string, passwordHash string) error
}

type cache struct {
	client cache2.Cache
}

func NewCache(client cache2.Cache) Cache {
	return &cache{
		client: client,
	}
}

func (c *cache) GetCode(ctx context.Context, email string) (string, error) {
	code, err := c.client.Get(ctx, email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("code is empty")
		}
		return "", err
	}
	val, ok := code.(string)
	if !ok {
		return "", errors.New("code is not string")
	}
	return val, nil
}

func (c *cache) SetCode(ctx context.Context, email string, code string) error {
	return c.client.SetWithExpiration(ctx, email, code, codeAndCredsLifeTime)
}

func (c *cache) GetPasswordHash(ctx context.Context, email string) (string, error) {
	pass, err := c.client.Get(ctx, email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("password hash is empty")
		}
		return "", err
	}
	val, ok := pass.(string)
	if !ok {
		return "", errors.New("password hash is not string")
	}
	return val, nil
}

func (c *cache) SetPasswordHash(ctx context.Context, email string, passwordHash string) error {
	return c.client.SetWithExpiration(ctx, email, passwordHash, codeAndCredsLifeTime)
}
