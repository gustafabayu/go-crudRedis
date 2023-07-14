package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gustafabayu/go-crudRedis/config"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	fmt.Println("connected successful to Redis")
	return rdb
}
