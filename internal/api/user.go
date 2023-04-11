package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/salmaan72/latitude-assignment/internal/user"

	"github.com/gin-gonic/gin"
)

func (api *API) Login(c *gin.Context) {
	login := user.Login{}
	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid format")
		return
	}

	ctx := c.Request.Context()
	user, err := api.UserService.ReadByEmail(ctx, login.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ver, err := api.VerifierService.ReadByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	hashed := api.VerifierService.PassHash([]byte(login.Password))
	if ver.Password != hashed {
		c.JSON(http.StatusUnauthorized, "invalid username/password")
		return
	}

	tokenDetails, err := api.AuthService.CreateToken(user.ID.String(), user.Username)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// save token ib redis
	err = api.AuthService.CreateAuthEntry(user.ID.String(), tokenDetails)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// set cookie
	cookie := http.Cookie{}
	cookie.Name = "accesstoken"
	cookie.Value = tokenDetails.AccessToken
	cookie.Expires = time.Now().Add(time.Second * 120)
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(c.Writer, &cookie)

	tokens := map[string]string{
		"access_token": tokenDetails.AccessToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func (api *API) MyInfo(c *gin.Context) {
	// accessDetails, err := api.AuthService.ExtractTokenMetadata(c.Request)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// userID, err := api.AuthService.FetchAuth(accessDetails.TokenUUID)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// ledger, err := api.LedgerService.Read(userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, ledger)
}

func (api *API) Signup(c *gin.Context) {
	ctx := c.Request.Context()

	user := &user.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		return
	}
	created, err := api.UserService.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, created)
}
