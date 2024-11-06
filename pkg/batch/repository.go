package batch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Database *redis.Client
	CTX      context.Context
}

func (client *RedisRepository) GetMessageQueue(userID string) error {
	rdb := client.Database
	ctx := client.CTX

	key := fmt.Sprintf("user:%s:documents", userID)
	vals, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}

	var documents []KafkaData
	for _, val := range vals {
		var doc KafkaData
		if err := json.Unmarshal([]byte(val), &doc); err != nil {
			return err
		}
		documents = append(documents, doc)
	}

	return nil
}
