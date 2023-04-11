package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salmaan72/latitude-assignment/internal/server/http/middleware"
)

type Route struct {
	Feature  string
	URI      string
	Method   string
	Handlers []gin.HandlerFunc
}

func (h *HTTP) routes() []Route {
	routes := []Route{
		{
			Feature: "user.user.r", // role.module.access
			URI:     "/myinfo",
			Method:  http.MethodGet,
			Handlers: []gin.HandlerFunc{
				h.API.MyInfo,
			},
		},
	}

	return routes
}

func (h *HTTP) InitializeRoutes() {
	// h.ginEngine.POST("/login", h.API.Login)
	h.ginEngine.POST("/auth/signup", h.API.Signup)
	router := h.ginEngine.Group("/")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// auth middleware
	router.Use(middleware.TokenAuthMiddleware())

	router.GET("/myinfo", h.API.MyInfo)
	router.GET("/dashboard", middleware.AuthoriseUserLedger(h.redisClient), h.API.FetchUserLedger)

}
