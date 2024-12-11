package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

var (
	connectionContextKey contextKey = "sdk.pgxpool.acquired.connection"
)

// AcquireConnAndSetIntoContext Acquires a connection from the pool and sets it into the context.
// Needed to use a single connection for multiple requests.
func AcquireConnAndSetIntoContext(ctx context.Context, pool *pgxpool.Pool) (context.Context, func(), error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to acquire connection: %w", err)
	}

	cleanup := func() {
		conn.Release()
	}

	return context.WithValue(ctx, connectionContextKey, conn), cleanup, nil
}

// ConnFromContext Return connection to db from context.
func ConnFromContext(ctx context.Context) *pgxpool.Conn {
	if conn, ok := ctx.Value(connectionContextKey).(*pgxpool.Conn); ok {
		return conn
	}
	return nil
}
