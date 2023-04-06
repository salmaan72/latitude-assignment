package http

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salmaan72/latitude-assignment/internal/api"
	"github.com/salmaan72/latitude-assignment/internal/auth"
	"github.com/salmaan72/latitude-assignment/internal/clients/cachestore"
)

type HTTP struct {
	config      *Config
	ginEngine   *gin.Engine
	redisClient *cachestore.Client
	Redis       auth.Auth
	TK          auth.Token
	API         *api.API
}

type Config struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func newGinEngine() (*gin.Engine, error) {
	ginEngine := gin.Default()

	return ginEngine, nil
}

func (h *HTTP) ListenAndServe() error {
	err := h.ginEngine.Run(fmt.Sprintf("%s:%s", h.config.Host, h.config.Port))
	if err != nil {
		return err
	}

	return nil
}

func NewHTTPServer(config *Config, redisClient *cachestore.Client, api *api.API) (*HTTP, error) {
	ginEngine, err := newGinEngine()
	if err != nil {
		return nil, err
	}

	return &HTTP{
		config:      config,
		ginEngine:   ginEngine,
		redisClient: redisClient,
		Redis:       api.AuthService,
		API:         api,
	}, nil
}
