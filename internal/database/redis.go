package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func GetRedisConnection(ctx context.Context) error {
	REDIS_PORT := os.Getenv("REDIS_PORT")
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	uri := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: REDIS_PASSWORD,
		DB:       0,
		Protocol: 2,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	log.Println("\033[32mCONNECTED TO REDIS DATABASE\033[0m")
	return nil
}
