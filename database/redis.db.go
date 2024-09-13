/*
 * @Author: lihuan
 * @Date: 2024-09-04 19:29:38
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-13 21:52:15
 * @Email: 17719495105@163.com
 */
package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/smartgreeting/mini-go/utils"
)

type RedisDB struct {
}

var (
	redisDB  *redis.Client
	redisCtx context.Context
)

func NewRedisDB(ctx context.Context, c *utils.Conf) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Dns,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	redisDB = rdb
	redisCtx = ctx
	return &RedisDB{}
}

func (*RedisDB) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redisDB.Set(redisCtx, key, value, expiration)

}

func (*RedisDB) Get(key string) *redis.StringCmd {
	return redisDB.Get(redisCtx, key)
}
func (*RedisDB) TTL(key string) *redis.DurationCmd {
	return redisDB.TTL(redisCtx, key)
}
func (*RedisDB) PTTL(key string) *redis.DurationCmd {
	return redisDB.PTTL(redisCtx, key)
}
