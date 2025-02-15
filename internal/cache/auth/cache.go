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

var (
	ErrSpamDetected = errors.New("spam detected")
)

type Cache interface {
	GetCode(ctx context.Context, email string) (string, error)
	SetCode(ctx context.Context, email string, code string) error
	CheckSpam(ctx context.Context, key string) (bool, error)
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
	key := c.genKey(codeKey, email)
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
	key := c.genKey(codeKey, email)
	spam, err := c.CheckSpam(ctx, key)
	if err != nil {
		return err
	}
	if spam {
		return ErrSpamDetected
	}
	return c.client.SetWithExpiration(ctx, key, code, codeLifeTime)
}

func (c *cache) GetPasswordHash(ctx context.Context, email string) (string, error) {
	key := c.genKey(passwordHashKey, email)
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
	key := c.genKey(passwordHashKey, email)
	return c.client.SetWithExpiration(ctx, key, passwordHash, credsLifeTime)
}

func (c *cache) CheckSpam(ctx context.Context, key string) (bool, error) {
	val, err := c.client.Exists(ctx, key)
	if err != nil {
		return false, err
	}
	return val, nil
}

func (c *cache) genKey(format, arg string) string {
	return fmt.Sprintf(format, arg)
}
