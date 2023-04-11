package controller

import (
	"net/http"

	"github.com/sebastianloose/strava-map-api/model"

	"github.com/sebastianloose/strava-map-api/service/cache"
	"github.com/sebastianloose/strava-map-api/service/strava"

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

	user, err := cache.GetUserById(userId)

	if err != nil {
		c.String(http.StatusBadRequest, "User not found")
		return
	}

	activities, err := strava.GetActivitiesForUser(user)

	var filteredActivities []model.Activity

	for _, a := range activities {
		if a.Map.SummaryPolyline != "" {
			filteredActivities = append(filteredActivities, a)
		}
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, filteredActivities)
}
