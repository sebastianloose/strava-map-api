package controller

import (
	"net/http"
	"strconv"

	"github.com/sebastianloose/strava-map-api/model"
	"github.com/sebastianloose/strava-map-api/model/strava_response"

	"github.com/sebastianloose/strava-map-api/service/strava"

	"github.com/gin-gonic/gin"
	"github.com/sebastianloose/strava-map-api/auth"
)

func GetActivities(c *gin.Context) {
	user, err := auth.ExtractUser(c)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	activities, err := strava.GetActivitiesOverviewForUser(user)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	var filteredActivities []strava_response.ActivitySummary

	for _, a := range activities {
		if a.Map.SummaryPolyline != "" {
			filteredActivities = append(filteredActivities, a)
		}
	}

	c.JSON(http.StatusOK, filteredActivities)
}

func GetActivity(c *gin.Context) {
	user, err := auth.ExtractUser(c)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	activityId, err := strconv.ParseInt(c.Param("activityId"), 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid activityId")
		return
	}

	details, err := strava.GetActivityForUser(user, activityId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	activity := model.ActivityDetailed{}

	activity.ActivityId = activityId
	activity.Points = details.Points.Data

	c.JSON(http.StatusOK, activity)
}
