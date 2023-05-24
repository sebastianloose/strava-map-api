package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sebastianloose/strava-map-api/model"
	"github.com/sebastianloose/strava-map-api/model/strava_response"
)

var baseUrl = "https://www.strava.com/api/v3"

func GetActivitiesOverview(user model.User) ([]strava_response.ActivitySummary, error) {
	requestUrl, _ := url.Parse(baseUrl + "/athlete/activities")

	query := requestUrl.Query()
	query.Add("per_page", "200")
	requestUrl.RawQuery = query.Encode()

	req, _ := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Set("Authorization", "Bearer "+user.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("strava activities overview request failed")
	}

	var activities []strava_response.ActivitySummary
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &activities)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

func GetActivity(user model.User, activityId int64) (strava_response.ActivityDetailed, error) {
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/activities/%d/streams", baseUrl, activityId))

	query := requestUrl.Query()
	query.Add("keys", "latlng")
	query.Add("key_by_type", "true")
	requestUrl.RawQuery = query.Encode()

	fmt.Println(requestUrl.String())

	req, _ := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Set("Authorization", "Bearer "+user.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return strava_response.ActivityDetailed{}, err
	}

	var activity strava_response.ActivityDetailed
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &activity)

	if resp.StatusCode != 200 {
		return strava_response.ActivityDetailed{}, errors.New("strava activity request failed")
	}

	return activity, nil
}
