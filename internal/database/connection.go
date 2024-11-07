package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type RedisConnection struct {
	Database *redis.Client
	CTX      context.Context
}

type MongoConnection struct {
	Database *mongo.Database
	CTX      context.Context
}
