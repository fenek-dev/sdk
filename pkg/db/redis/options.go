package redis

import (
	"crypto/tls"

	"github.com/redis/go-redis/v9"
)

func WithTLS(tlsConfig *tls.Config) Option {
	return func(opt *redis.Options) {
		opt.TLSConfig = tlsConfig
	}
}

func WithMinIdleConnections(minIdleConnections int) Option {
	return func(opt *redis.Options) {
		opt.MinIdleConns = minIdleConnections
	}
}
