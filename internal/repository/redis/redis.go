package goredis

import (
	"github.com/Levap123/adverts/configs"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	cl *redis.Client
}

func NewRedis(confs configs.RedisConf) *RedisClient {
	cl := redis.NewClient(&redis.Options{
		Addr:     confs.Host,
		Password: confs.Password,
		DB:       confs.DB,
	})
	return &RedisClient{
		cl: cl,
	}
}
