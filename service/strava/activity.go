package strava

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sebastianloose/strava-map-api/model"
)

func GetActivitiesForUser(user model.User) ([]model.Activity, error) {
	url, _ := url.Parse("https://www.strava.com/api/v3/athlete/activities")

	query := url.Query()
	query.Add("per_page", "200")
	url.RawQuery = query.Encode()

	req, _ := http.NewRequest(http.MethodGet, url.String(), nil)
	req.Header.Set("Authorization", "Bearer "+user.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("strava activity request failed")
	}

	var activities []model.Activity
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &activities)

	if err != nil {
		return nil, err
	}

	return activities, nil
}
