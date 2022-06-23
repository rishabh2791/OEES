package persistance

import (
	"oees/infrastructure/utilities"

	"github.com/go-redis/redis"
)

type RedisService struct {
	RedisClient *redis.Client
}

func NewCacheStore(config utilities.ServerConfig) (*RedisService, error) {
	cacheConfig := config.GetCacheConfig()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.CacheHost + ":" + cacheConfig.CachePort,
		Password: cacheConfig.CachePassword,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &RedisService{
		RedisClient: redisClient,
	}, nil
}
