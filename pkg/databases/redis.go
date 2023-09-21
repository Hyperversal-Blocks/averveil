package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Set(ctx context.Context, key string, values interface{}) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}

type redisClient struct {
	client *redis.Client
}

func NewClient(addr string, db int) Redis {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
	if client == nil {
		panic("failed to connect to redis")
	}

	return &redisClient{
		client: client,
	}
}

func (r *redisClient) Set(ctx context.Context, key string, values interface{}) error {
	err := r.client.Set(ctx, key, values, time.Minute*2).Err()
	if err != nil {
		return fmt.Errorf("error setting key in redis: %w", err)
	}
	return nil
}

func (r *redisClient) Get(ctx context.Context, key string) interface{} {
	return r.client.Get(ctx, key)
}

func (r *redisClient) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error removing key in redis: %w", err)
	}
	return nil
}
