package postgres

// import (
// 	"context"

// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/jackc/pgx/v5/pgxtrace"
// )

// func NewPostgresPool(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
// 	config, err := pgxpool.ParseConfig(connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	config.AfterConnect = pgxtrace.AfterConnect
// 	pool, err := pgxpool.ConnectConfig(ctx, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pool, nil
// }
