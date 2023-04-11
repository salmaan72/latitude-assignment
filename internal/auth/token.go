package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Token interface {
	CreateToken(userID, username string) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

func (s *Service) CreateToken(userID, username string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.ExpiresAt = time.Now().Add(time.Second * 100).Unix() //expires after 100 sec

	td.TokenUUID = uuid.New().String()

	var err error
	//Creating AccessDetails Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUUID
	atClaims["user_id"] = userID
	atClaims["user_name"] = username
	atClaims["exp"] = td.ExpiresAt
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}
func (s *Service) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func IsTokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	// bearToken := r.Header.Get("Authorization")
	// strArr := strings.Split(bearToken, " ")
	// if len(strArr) == 2 {
	// 	return strArr[0]
	// }
	cookie, err := r.Cookie("accesstoken")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func Extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		userID, userOk := claims["user_id"].(string)
		username, userNameOk := claims["user_name"].(string)
		if !ok || !userOk || !userNameOk {
			return nil, errors.New("unauthorized")
		} else {
			return &AccessDetails{
				TokenUUID: accessUUID,
				UserID:    userID,
				Username:  username,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
