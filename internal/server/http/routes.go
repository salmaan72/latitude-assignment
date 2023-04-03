package http

import (
	"github.com/gin-gonic/gin"
	"github.com/salmaan72/latitude-assignment/internal/server/http/middleware"
)

func (h *HTTP) InitializeRoutes() {
	h.ginEngine.POST("/login", h.API.Login)

	router := h.ginEngine.Group("/")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// auth middleware
	router.Use(middleware.TokenAuthMiddleware())

	router.GET("/myinfo", h.API.MyInfo)
	router.GET("/dashboard", middleware.AuthoriseUserLedger(h.redisClient), h.API.FetchUserLedger)
}
