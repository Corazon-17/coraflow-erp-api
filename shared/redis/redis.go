package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	RDB *redis.Client
}

func NewRedis(addr string) *Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &Client{RDB: rdb}
}

func (c *Client) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return c.RDB.Set(ctx, key, value, ttl).Err()
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.RDB.Get(ctx, key).Result()
}

func (c *Client) Delete(ctx context.Context, key string) error {
	return c.RDB.Del(ctx, key).Err()
}
