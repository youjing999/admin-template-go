package redis

import (
	"admin-template-go/common/config"
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Address,
		DB:   3,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
