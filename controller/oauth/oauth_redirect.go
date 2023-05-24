package oauth

import (
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastianloose/strava-map-api/auth"
	"github.com/sebastianloose/strava-map-api/model"
	"github.com/sebastianloose/strava-map-api/service/cache"
	"github.com/sebastianloose/strava-map-api/service/strava"
)

func Redirect(c *gin.Context) {
	if len(c.Request.URL.Query()["code"]) < 1 {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	code := c.Request.URL.Query()["code"][0]

	userToken, err := strava.TokenExchange(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	var user model.User
	user.AccessToken = userToken.AccessToken
	user.ExpiresAt = userToken.ExpiresAt
	user.UserId = userToken.StravaAccount.Id

	token, err := auth.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	cache.AddUser(user)

	redirectUrl, _ := url.Parse(os.Getenv("APP_URL"))

	query := redirectUrl.Query()
	query.Add("token", token)
	query.Add("expiresAt", strconv.Itoa(user.ExpiresAt))
	redirectUrl.RawQuery = query.Encode()

	c.Redirect(http.StatusTemporaryRedirect, redirectUrl.String())
}
