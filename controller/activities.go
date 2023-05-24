package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sebastianloose/strava-map-api/model"

	"github.com/sebastianloose/strava-map-api/service/strava"

	"github.com/gin-gonic/gin"
	"github.com/sebastianloose/strava-map-api/auth"
	"github.com/twpayne/go-polyline"
)

func GetActivities(c *gin.Context) {
	user, err := auth.ExtractUser(c)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	rawActivities, err := strava.GetActivitiesOverview(user)

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	var activities []model.Activity

	for _, a := range rawActivities {
		if a.Map.SummaryPolyline == "" {
			continue
		}

		activity := model.Activity{
			Id:                 a.Id,
			Name:               a.Name,
			Distance:           a.Distance,
			MovingTime:         a.MovingTime,
			ElapsedTime:        a.ElapsedTime,
			StartDate:          a.StartDate,
			TotalElevationGain: a.TotalElevationGain,
			Type:               a.Type,
			PolylineRoute:      a.Map.SummaryPolyline,
			IsDetailedRoute:    false,
			AverageSpeed:       a.AverageSpeed,
			MaximumSpeed:       a.MaximumSpeed,
			AverageHeartRate:   a.AverageHeartRate,
			MaximumHeartRate:   a.MaximumHeartRate,
			AverageCadence:     a.AverageCadence,
			AverageTemperature: a.AverageTemperature,
			AverageWatts:       a.AveragePower,
		}

		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, activities)
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

	details, err := strava.GetActivity(user, activityId)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	route := string(polyline.EncodeCoords(details.Points.Data))

	activity := model.ActivityDetailedRoute{}

	activity.ActivityId = activityId
	activity.PolylineRoute = route

	c.JSON(http.StatusOK, activity)
}
