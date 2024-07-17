package util

import (
	"admin-template-go/common/constant"
	"admin-template-go/pkg/redis"
	"context"
	"log"
	"time"
)

var ctx = context.Background()

type RedisStore struct{}

// Set 存验证码
func (r RedisStore) Set(id string, value string) {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
	}
}

// Get 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Verify 检验验证码
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
