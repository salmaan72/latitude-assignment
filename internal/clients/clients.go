package clients

import (
	"github.com/go-redis/redis"
	localredis "github.com/salmaan72/latitude-assignment/internal/clients/redis"
)

type Service struct {
	RedisClient *redis.Client
}

func NewService(redisCfg *localredis.Config) *Service {
	return &Service{
		RedisClient: localredis.NewService(redisCfg).RedisClient,
	}
}
