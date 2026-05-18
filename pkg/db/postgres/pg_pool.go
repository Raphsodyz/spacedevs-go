package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Raphsodyz/spacedevs-go/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cnfgName string) (*pgxpool.Pool, error) {
	viper, err := config.LoadConfig(cnfgName)
	if err != nil {
		return nil, err
	}

	cnfg, err := config.ParseConfig(viper)
	if err != nil {
		return nil, err
	}

	pg := cnfg.Postgresql
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pg.PostgresqlUser,
		pg.PostgresqlPassword,
		pg.PostgresqlHost,
		pg.PostgresqlPort,
		pg.PostgresqlDbname,
		boolToSSLMode(pg.PostgresqlSSLMode),
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
