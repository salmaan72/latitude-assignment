package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/salmaan72/latitude-assignment/internal/auth"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.IsTokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthoriseUserLedger(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Request.URL.Query().Get("userID")

		userAccessToken, err := redisClient.Get(userID).Result()
		if err != nil {
			c.Next()
			return
		}

		_, err = jwt.Parse(userAccessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("Token is expired")) {
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.JSON(http.StatusUnauthorized, "resource already in use")
		c.Abort()
	}
}
