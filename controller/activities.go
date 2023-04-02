package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebastianloose/strava-map-api/auth"
)

func GetActivities(c *gin.Context) {
	token := auth.GetTokenFromRequest(c)
	userId, err := auth.ExtractUserID(token)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}
