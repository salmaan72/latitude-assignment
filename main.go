package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/salmaan72/latitude-assignment/internal/api"
	"github.com/salmaan72/latitude-assignment/internal/auth"
	redisService "github.com/salmaan72/latitude-assignment/internal/clients/redis"
	"github.com/salmaan72/latitude-assignment/internal/configs"
	"github.com/salmaan72/latitude-assignment/internal/ledger"
	"github.com/salmaan72/latitude-assignment/internal/server"
	"github.com/salmaan72/latitude-assignment/internal/server/http"
	"github.com/salmaan72/latitude-assignment/internal/user"
)

func initAPIService(cfgHandler *configs.Config, redisClient *redis.Client) (*api.API, error) {
	userService, err := user.NewService()
	if err != nil {
		return nil, err
	}

	authService := auth.NewAuthService(redisClient)

	ledgerService := ledger.NewService()

	apiService := api.New(userService, authService, ledgerService)

	return apiService, nil

}

func main() {
	cfgHandler, err := configs.NewConfigHandler(".env")
	if err != nil {
		log.Fatal(err)
	}
	redisService := redisService.NewService(cfgHandler.Redis())

	apiService, err := initAPIService(cfgHandler, redisService.RedisClient)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := http.NewHTTPServer(cfgHandler.HTTP(), redisService.RedisClient, apiService)
	if err != nil {
		log.Fatal(err)
	}
	httpServer.InitializeRoutes()
	err = server.StartServer(httpServer)
	if err != nil {
		log.Fatal(err)
	}
}
