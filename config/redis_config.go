package config

import (
	goredis "github.com/go-redis/redis/v8"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/isyscore/isc-gobase/redis"
)

var RedisDb goredis.UniversalClient

func init() {
	RedisDbTem, err := redis.GetClient()
	if err != nil {
		logger.Warn("redis连接异常", err)
	}
	RedisDb = RedisDbTem
}
