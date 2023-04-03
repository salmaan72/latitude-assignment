package redis

import "github.com/go-redis/redis"

type Service struct {
	RedisClient *redis.Client
}

type Config struct {
	Host     string
	Port     string
	Password string
}

func newRedisDB(config *Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       0,
	})

	return redisClient
}

func NewService(config *Config) *Service {
	return &Service{
		RedisClient: newRedisDB(config),
	}
}
