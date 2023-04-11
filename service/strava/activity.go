package strava

import (
	"encoding/json"
	"errors"
	"github.com/sebastianloose/strava-map-api/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetActivitiesForUser(user model.User) ([]model.Activity, error) {
	tokenUrl, err := url.Parse("https://www.strava.com/api/v3/athlete/activities")

	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodGet, tokenUrl.String(), nil)
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
