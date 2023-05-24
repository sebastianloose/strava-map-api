package model

import (
	"time"
)

type Activity struct {
	Id                 int64     `json:"id"`
	Name               string    `json:"name"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"movingTime"`
	ElapsedTime        int       `json:"elapsedTime"`
	StartDate          time.Time `json:"startDate"`
	TotalElevationGain float64   `json:"totalElevationGain"`
	Type               string    `json:"type"`
	PolylineRoute      string    `json:"polylineRoute"`
	IsDetailedRoute    bool      `json:"isDetailedRoute"`
	AverageSpeed       float64   `json:"averageSpeed"`
	MaximumSpeed       float64   `json:"maxSpeed"`
	AverageHeartRate   float64   `json:"averageHeartRate"`
	MaximumHeartRate   float64   `json:"maxHeartRate"`
	AverageCadence     float64   `json:"averageCadence"`
	AverageTemperature float64   `json:"averageTemp"`
	AverageWatts       float64   `json:"averageWatts"`
}

type ActivityDetailedRoute struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`
	PolylineRoute string `json:"polylineRoute"`
}
