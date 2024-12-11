package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host string `env:"HOST" env-required:"true"`
	Port string `env:"PORT" env-required:"true"`
	Name string `env:"NAME" env-required:"true"`
	User string `env:"USER" env-required:"true"`
	Pass string `env:"PASS" env-required:"true"`
	SSL  string `env:"SSL"`
}

type Option func(config *pgxpool.Config) *pgxpool.Config

func Connect(ctx context.Context, args Config, options ...Option) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		args.User,
		args.Pass,
		args.Host,
		args.Port,
		args.Name,
		args.SSL,
	)

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database conf: %w", err)
	}
	for _, option := range options {
		conf = option(conf)
	}

	conn, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return conn, nil
}

func MustConnect(ctx context.Context, args Config, options ...Option) *pgxpool.Pool {
	conn, err := Connect(ctx, args, options...)
	if err != nil {
		panic(err)
	}

	return conn
}

func Ping(ctx context.Context, conn *pgxpool.Pool, timeout time.Duration) error {
	pingCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := conn.Ping(pingCtx)
	if err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	return nil
}

func MustPing(ctx context.Context, conn *pgxpool.Pool, timeout time.Duration) {
	if err := Ping(ctx, conn, timeout); err != nil {
		panic(err)
	}
}
