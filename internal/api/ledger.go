package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (api *API) FetchUserLedger(c *gin.Context) {
	userID := c.Request.URL.Query().Get("userID")
	ctx := c.Request.Context()
	id, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, "invalid user id")
		c.Abort()
		return
	}
	l, err := api.LedgerService.ReadByUserID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "resource not found")
		c.Abort()
		return
	}

	// c.JSON(http.StatusOK, l)
	c.AbortWithStatusJSON(http.StatusOK, l)
}
