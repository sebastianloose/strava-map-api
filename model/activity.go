package model

import "github.com/sebastianloose/strava-map-api/model/strava_response"

type ActivityDetailed struct {
	ActivityId int64                      `json:"activityId"`
	Points     []strava_response.Location `json:"points"`
}
