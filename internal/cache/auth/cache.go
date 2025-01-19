package auth

import (
	"context"
	"errors"
	"fmt"
	cache2 "github.com/Rasikrr/learning_platform/internal/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	codeLifeTime    = time.Minute * 5
	credsLifeTime   = time.Minute * 6
	codeKey         = "code:%s"
	passwordHashKey = "password:%s"
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
	key := c.genCodeKey(email)
	code, err := c.client.Get(ctx, key)
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
	key := c.genCodeKey(email)
	return c.client.SetWithExpiration(ctx, key, code, codeLifeTime)
}

func (c *cache) GetPasswordHash(ctx context.Context, email string) (string, error) {
	key := c.genPasswordKey(email)
	pass, err := c.client.Get(ctx, key)
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
	key := c.genPasswordKey(email)
	return c.client.SetWithExpiration(ctx, key, passwordHash, credsLifeTime)
}

func (c *cache) genCodeKey(email string) string {
	return fmt.Sprintf(codeKey, email)
}

func (c *cache) genPasswordKey(email string) string {
	return fmt.Sprintf(passwordHashKey, email)
}
