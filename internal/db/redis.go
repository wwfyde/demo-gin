package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisClient struct {
	client *redis.Client
	cfg    *viper.Viper
}

func NewRedisClient(client *redis.Client, cfg *viper.Viper) *RedisClient {
	return &RedisClient{
		client: client,
		cfg:    cfg,
	}
}
