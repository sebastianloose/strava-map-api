package strava_response

type StravaAccount struct {
	Id int64 `json:"id"`
}

type Auth struct {
	AccessToken   string        `json:"access_token"`
	ExpiresAt     int           `json:"expires_at"`
	StravaAccount StravaAccount `json:"athlete"`
}
