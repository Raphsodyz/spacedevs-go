package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Raphsodyz/spacedevs-go/config"
	"github.com/redis/go-redis/v9"
)

const (
	maxRetries      = 3
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 1 * time.Second
	dialTimeout     = 5 * time.Second
	readTimeout     = 3 * time.Second
	writeTimeout    = 3 * time.Second
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            cfg.Redis.RedisAddr,
		Password:        cfg.Redis.Password,
		DB:              cfg.Redis.DB,
		MinIdleConns:    cfg.Redis.MinIdleConns,
		PoolSize:        cfg.Redis.PoolSize,
		PoolTimeout:     time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis: failed to connect to %s: %w", cfg.Redis.RedisAddr, err)
	}

	return client, nil
}
