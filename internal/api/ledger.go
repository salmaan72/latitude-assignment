package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) FetchUserLedger(c *gin.Context) {
	userID := c.Request.URL.Query().Get("userID")
	l, err := api.LedgerService.Read(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, "resource not found")
		c.Abort()
		return
	}

	// c.JSON(http.StatusOK, l)
	c.AbortWithStatusJSON(http.StatusOK, l)
}
