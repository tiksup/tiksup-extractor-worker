package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	Database *redis.Client
	CTX      context.Context
}
