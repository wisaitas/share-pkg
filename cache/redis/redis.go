package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Util interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (bool, error)
}

type util struct {
	Client *redis.Client
}

func New(client *redis.Client) Util {
	return &util{
		Client: client,
	}
}

func (r *util) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *util) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *util) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *util) Exists(ctx context.Context, keys ...string) (bool, error) {
	exists, err := r.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
