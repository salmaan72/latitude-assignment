package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	redisclient "github.com/salmaan72/latitude-assignment/internal/clients/redis"
	"github.com/salmaan72/latitude-assignment/internal/server/http"
)

type Config struct{}

func (cfg *Config) HTTP() *http.Config {
	return &http.Config{
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
}

func (cfg *Config) Redis() *redisclient.Config {
	return &redisclient.Config{
		Host:     os.Getenv("HOST_REDIS"),
		Port:     os.Getenv("PORT_REDIS"),
		Password: os.Getenv("PASS_REDIS"),
	}
}

func NewConfigHandler(envFilePath string) (*Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		return nil, err
	}

	return &Config{}, nil
}
