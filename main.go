package main

import (
	"log"

	"github.com/salmaan72/latitude-assignment/internal/api"
	"github.com/salmaan72/latitude-assignment/internal/auth"
	"github.com/salmaan72/latitude-assignment/internal/clients"
	"github.com/salmaan72/latitude-assignment/internal/configs"
	"github.com/salmaan72/latitude-assignment/internal/server"
	"github.com/salmaan72/latitude-assignment/internal/server/http"
	"github.com/salmaan72/latitude-assignment/internal/user"
	"github.com/salmaan72/latitude-assignment/internal/user/ledger"
)

func initAPIService(cfgHandler *configs.Config, clients *clients.Service) (*api.API, error) {
	ledgerService, err := ledger.NewService(clients.DatastoreClient)
	if err != nil {
		return nil, err
	}

	userService, err := user.NewService(clients.DatastoreClient, ledgerService)
	if err != nil {
		return nil, err
	}

	authService := auth.NewAuthService(clients.RedisClient)

	apiService := api.New(userService, authService, ledgerService)

	return apiService, nil

}

func main() {
	cfgHandler, err := configs.NewConfigHandler(".env")
	if err != nil {
		log.Fatal(err)
	}
	clients, err := clients.NewService(cfgHandler.Redis(), cfgHandler.Datastore())
	if err != nil {
		log.Fatal(err)
	}

	// postgres, err := pos

	apiService, err := initAPIService(cfgHandler, clients)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := http.NewHTTPServer(cfgHandler.HTTP(), clients.RedisClient, apiService)
	if err != nil {
		log.Fatal(err)
	}
	httpServer.InitializeRoutes()
	err = server.StartServer(httpServer)
	if err != nil {
		log.Fatal(err)
	}
}
