package redis

import (
	"context"
	"fmt"
	"golang-project-template/config"

	redis "github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	rdb    *redis.Client
	expire time.Duration
}

type Val interface{}

func NewClient(cfg config.Config) (*Client, error) {
	dur, _ := time.ParseDuration(cfg.Redis.ExpireTime)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Client{
		rdb:    rdb,
		expire: dur,
	}, nil
}

func (c *Client) Set(ctx context.Context, key string, value Val) error {
	if err := c.rdb.Set(ctx, key, value, c.expire).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return val, err
	} else if err != nil {
		return "", err
	}
	return val, err
}
