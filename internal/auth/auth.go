package auth

import (
	"errors"

	"github.com/salmaan72/latitude-assignment/internal/clients/cachestore"
)

type AccessDetails struct {
	TokenUUID string
	UserID    string
	Username  string
}

type TokenDetails struct {
	TokenUUID   string
	AccessToken string
	ExpiresAt   int64
}

type Auth interface {
	CreateAuthEntry(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteTokens(*AccessDetails) error
}

type Service struct {
	redisClient *cachestore.Client
}

// Save token metadata to Redis
func (s *Service) CreateAuthEntry(userID string, td *TokenDetails) error {
	record, err := s.redisClient.Set(td.TokenUUID, userID, 0).Result()
	if err != nil {
		return err
	}
	tokenRecord, err := s.redisClient.Set(userID, td.AccessToken, 0).Result()
	if err != nil {
		return err
	}

	if record == "0" || tokenRecord == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

func (s *Service) FetchAuth(tokenUUID string) (string, error) {
	userID, err := s.redisClient.Get(tokenUUID).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *Service) DeleteTokens(ad *AccessDetails) error {
	deletedAt, err := s.redisClient.Del(ad.TokenUUID).Result()
	if err != nil {
		return err
	}

	if deletedAt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}

func NewAuthService(client *cachestore.Client) *Service {
	return &Service{redisClient: client}
}
