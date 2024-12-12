package pg

import "github.com/jackc/pgx/v5/pgxpool"

func WithMaxConnections(maxConnections int32) Option {
	return func(config *pgxpool.Config) *pgxpool.Config {
		config.MaxConns = maxConnections
		return config
	}
}

func WithMinConnections(minConnections int32) Option {
	return func(config *pgxpool.Config) *pgxpool.Config {
		config.MinConns = minConnections
		return config
	}
}
