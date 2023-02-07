package goredis

import (
	"github.com/Levap123/adverts/configs"
	"github.com/redis/go-redis/v9"
)

func InitRedis(confs configs.RedisConf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     confs.Host,
		Password: confs.Password,
		DB:       confs.DB,
	})
	return client
}
