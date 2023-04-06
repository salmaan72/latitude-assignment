package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/salmaan72/latitude-assignment/internal/clients/cachestore"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
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

func (cfg *Config) Redis() *cachestore.Config {
	return &cachestore.Config{
		Host:     os.Getenv("HOST_REDIS"),
		Port:     os.Getenv("PORT_REDIS"),
		Password: os.Getenv("PASS_REDIS"),
	}
}

func (cfg *Config) Datastore() *datastore.Config {
	return &datastore.Config{
		Host:         os.Getenv("HOST_DATASTORE"),
		Port:         os.Getenv("PORT_DATASTORE"),
		User:         os.Getenv("USER_DATASTORE"),
		Password:     os.Getenv("PASS_DATASTORE"),
		DBname:       os.Getenv("DBNAME_DATASTORE"),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
}

func NewConfigHandler(envFilePath string) (*Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		return nil, err
	}

	return &Config{}, nil
}
