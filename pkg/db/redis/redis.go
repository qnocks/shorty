package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type Config struct {
	Host     string
	Port     string
	Password string
}

type Client struct {
	Redis *redis.Client
	Ctx   context.Context
}

func NewRedis(cfg Config) *Client {
	return &Client{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Password: cfg.Password,
			DB:       0,
		}), ctx,
	}
}
