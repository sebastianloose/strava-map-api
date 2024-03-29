package strava_response

import (
	"time"
)

type Map struct {
	Id              string `json:"id"`
	SummaryPolyline string `json:"summary_polyline"`
}

type ActivitySummary struct {
	Id                   int64     `json:"id"`
	ExternalId           string    `json:"external_id"`
	UploadId             int64     `json:"upload_id"`
	Name                 string    `json:"name"`
	Distance             float64   `json:"distance"`
	MovingTime           int       `json:"moving_time"`
	ElapsedTime          int       `json:"elapsed_time"`
	TotalElevationGain   float64   `json:"total_elevation_gain"`
	Type                 string    `json:"type"`
	StartDate            time.Time `json:"start_date"`
	StartDateLocal       time.Time `json:"start_date_local"`
	TimeZone             string    `json:"time_zone"`
	StartLocation        []float64 `json:"start_latlng"`
	EndLocation          []float64 `json:"end_latlng"`
	City                 string    `json:"location_city"`
	State                string    `json:"location_state"`
	Country              string    `json:"location_country"`
	AchievementCount     int       `json:"achievement_count"`
	KudosCount           int       `json:"kudos_count"`
	CommentCount         int       `json:"comment_count"`
	AthleteCount         int       `json:"athlete_count"`
	PhotoCount           int       `json:"photo_count"`
	Map                  Map       `json:"map"`
	Trainer              bool      `json:"trainer"`
	Commute              bool      `json:"commute"`
	Manual               bool      `json:"manual"`
	Private              bool      `json:"private"`
	Flagged              bool      `json:"flagged"`
	GearId               string    `json:"gear_id"`
	AverageSpeed         float64   `json:"average_speed"`
	MaximumSpeed         float64   `json:"max_speed"`
	AverageCadence       float64   `json:"average_cadence"`
	AverageTemperature   float64   `json:"average_temp"`
	AveragePower         float64   `json:"average_watts"`
	WeightedAveragePower int       `json:"weighted_average_watts"`
	Kilojoules           float64   `json:"kilojoules"`
	DeviceWatts          bool      `json:"device_watts"`
	AverageHeartRate     float64   `json:"average_heartrate"`
	MaximumHeartRate     float64   `json:"max_heartrate"`
	Truncated            int       `json:"truncated"`
	HasKudoed            bool      `json:"has_kudoed"`
}

type PointsList struct {
	Data [][]float64 `json:"data"`
}

type ActivityDetailed struct {
	Points PointsList `json:"latlng"`
}
