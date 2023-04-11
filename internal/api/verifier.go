package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (api *API) Verify(c *gin.Context) {
	ctx := c.Request.Context()

	accessDetails, err := api.AuthService.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := api.AuthService.FetchAuth(accessDetails.TokenUUID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	comType := c.Request.URL.Query().Get("type")
	otp := c.Request.URL.Query().Get("otp")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	u, err := api.UserService.VerifyUser(ctx, parsedUserID, comType, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, u)
}
