package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Raphsodyz/spacedevs-go/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(pg config.PostgresConfig) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pg.User,
		pg.Password,
		pg.Host,
		pg.Port,
		pg.Db,
		boolToSSLMode(pg.Sslmode),
	)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func boolToSSLMode(enabled bool) string {
	if enabled {
		return "require"
	}
	return "disable"
}
