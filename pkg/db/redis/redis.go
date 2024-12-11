package redis

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host       string `env-required:"true" env:"HOST"`
	Port       string `env-required:"true" env:"PORT"`
	Pass       string `env-required:"true" env:"PASS"`
	DB         int    `env-required:"true" env:"DB"`
	User       string `env-required:"true" env:"USER"`
	ClientName string `env-required:"true" env:"CLIENT_NAME"`
}

type Option func(*redis.Options)

func New(cfg *Config, options ...Option) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:       net.JoinHostPort(cfg.Host, cfg.Port),
		Username:   cfg.User,
		Password:   cfg.Pass,
		DB:         cfg.DB,
		ClientName: cfg.ClientName,
	}

	for _, option := range options {
		option(opt)
	}

	rdb := redis.NewClient(opt)

	return rdb, nil
}

func MustNew(cfg Config, options ...Option) *redis.Client {
	rdb, err := New(&cfg, options...)
	if err != nil {
		panic(err)
	}
	return rdb
}

func Ping(ctx context.Context, conn *redis.Client, timeout time.Duration) error {
	pingCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := conn.Ping(pingCtx).Err()
	if err != nil {
		return fmt.Errorf("unable to ping redis: %w", err)
	}

	return nil
}

func MustPing(ctx context.Context, conn *redis.Client, timeout time.Duration) {
	err := Ping(ctx, conn, timeout)
	if err != nil {
		panic(err)
	}
}
