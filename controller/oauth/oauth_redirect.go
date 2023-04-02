package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	user.UserId, err = uuid.NewRandom()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	token, err := auth.GenerateToken(user.UserId, user.ExpiresAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	cache.User = append(cache.User, user)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
