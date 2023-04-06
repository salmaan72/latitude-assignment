package datastore

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	*gorm.DB
}

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBname       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewClient(config *Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBname,
		config.Port,
	)

	datastoreClient, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Client{datastoreClient}, nil
}
