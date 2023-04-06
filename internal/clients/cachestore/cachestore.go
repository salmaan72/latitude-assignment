package cachestore

import (
	"log"

	"github.com/go-redis/redis"
)

type Client struct {
	*redis.Client
}

type Config struct {
	Host     string
	Port     string
	Password string
}

func newRedisDB(config *Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       0,
	})

	resp, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Print(resp)

	return redisClient, nil
}

func NewClient(config *Config) (*Client, error) {
	client, err := newRedisDB(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
