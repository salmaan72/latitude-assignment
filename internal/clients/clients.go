package clients

import (
	"github.com/salmaan72/latitude-assignment/internal/clients/cachestore"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type Service struct {
	RedisClient     *cachestore.Client
	DatastoreClient *datastore.Client
}

func NewService(redisCfg *cachestore.Config, datastoreCfg *datastore.Config) (*Service, error) {
	redisClient, err := cachestore.NewClient(redisCfg)
	if err != nil {
		return nil, err
	}

	datastoreClient, err := datastore.NewClient(datastoreCfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		RedisClient:     redisClient,
		DatastoreClient: datastoreClient,
	}, nil
}
