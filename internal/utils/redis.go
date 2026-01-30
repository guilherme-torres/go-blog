package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{client: client}
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, exp time.Duration) error {
	if err := r.client.Set(ctx, key, value, exp).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}
