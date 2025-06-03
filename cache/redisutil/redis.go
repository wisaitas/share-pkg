package redisutil

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUtil interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (bool, error)
}

type redisUtil struct {
	Client *redis.Client
}

func NewRedisUtil(client *redis.Client) RedisUtil {
	return &redisUtil{
		Client: client,
	}
}

func (r *redisUtil) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := r.Client.Set(ctx, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("[Share Package RedisUtil] : %w", err)
	}

	return nil
}

func (r *redisUtil) Get(ctx context.Context, key string) (string, error) {
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("[Share Package RedisUtil] : %w", err)
	}

	return value, nil
}

func (r *redisUtil) Del(ctx context.Context, keys ...string) error {
	if err := r.Client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("[Share Package RedisUtil] : %w", err)
	}

	return nil
}

func (r *redisUtil) Exists(ctx context.Context, keys ...string) (bool, error) {
	exists, err := r.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, fmt.Errorf("[Share Package RedisUtil] : %w", err)
	}

	return exists > 0, nil
}
