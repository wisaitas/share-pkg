package redis

import (
	"context"
	"fmt"
	"time"

	redisLib "github.com/redis/go-redis/v9"
)

type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (bool, error)
}

type redis struct {
	Client *redisLib.Client
}

func NewRedis(client *redisLib.Client) Redis {
	return &redis{
		Client: client,
	}
}

func (r *redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := r.Client.Set(ctx, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("[redis] : %w", err)
	}

	return nil
}

func (r *redis) Get(ctx context.Context, key string) (string, error) {
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("[redis] : %w", err)
	}

	return value, nil
}

func (r *redis) Del(ctx context.Context, keys ...string) error {
	if err := r.Client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("[redis] : %w", err)
	}

	return nil
}

func (r *redis) Exists(ctx context.Context, keys ...string) (bool, error) {
	exists, err := r.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, fmt.Errorf("[redis] : %w", err)
	}

	return exists > 0, nil
}
