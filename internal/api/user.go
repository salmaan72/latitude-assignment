package api

import (
	"encoding/json"
	"net/http"

	"github.com/salmaan72/latitude-assignment/internal/user"

	"github.com/gin-gonic/gin"
)

// func (api *API) Login(c *gin.Context) {
// 	login := user.User{}
// 	err := json.NewDecoder(c.Request.Body).Decode(&login)
// 	if err != nil {
// 		return
// 	}

// 	//find user with username
// 	user, err := api.UserService.ReadByUsername(login.Username)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
// 		return
// 	}
// 	//compare the user from the request, with the one we defined:
// 	if user.Username != login.Username || user.Password != login.Password {
// 		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
// 		return
// 	}
// 	tokenDetails, err := api.AuthService.CreateToken(user.ID, user.Username)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}

// 	// save token ib redis
// 	err = api.AuthService.CreateAuthEntry(user.ID, tokenDetails)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	// set cookie
// 	cookie := http.Cookie{}
// 	cookie.Name = "accesstoken"
// 	cookie.Value = tokenDetails.AccessToken
// 	cookie.Expires = time.Now().Add(time.Second * 120)
// 	cookie.HttpOnly = true
// 	cookie.Path = "/"
// 	http.SetCookie(c.Writer, &cookie)

// 	tokens := map[string]string{
// 		"access_token": tokenDetails.AccessToken,
// 	}
// 	c.JSON(http.StatusOK, tokens)
// }

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
