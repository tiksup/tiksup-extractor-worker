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

func (client *RedisRepository) GetMessageQueue(userID string) ([]UserInfo, error) {
	rdb := client.Database
	ctx := client.CTX

	key := fmt.Sprintf("user:%s:documents", userID)
	vals, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	documents := []UserInfo{}
	for _, val := range vals {
		var doc UserInfo
		if err := json.Unmarshal([]byte(val), &doc); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}
