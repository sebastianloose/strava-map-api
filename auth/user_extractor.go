package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sebastianloose/strava-map-api/model"
	"github.com/sebastianloose/strava-map-api/service/cache"
)

func ExtractUser(c *gin.Context) (model.User, error) {
	token := GetTokenFromRequest(c)
	userId, err := ExtractUserID(token)

	if err != nil {
		return model.User{}, errors.New("invalid token")
	}

	user, err := cache.GetUser(userId)

	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}
