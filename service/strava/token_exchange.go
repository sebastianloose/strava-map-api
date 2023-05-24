package strava

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/sebastianloose/strava-map-api/model/strava_response"
)

func TokenExchange(userCode string) (strava_response.Auth, error) {
	tokenUrl, err := url.Parse("https://www.strava.com/oauth/token")

	if err != nil {
		log.Fatal(err)
	}

	queryParams := tokenUrl.Query()
	queryParams.Add("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	queryParams.Add("client_secret", os.Getenv("STRAVA_CLIENT_SECRET"))
	queryParams.Add("code", userCode)
	queryParams.Add("grant_type", "authorization_code")
	tokenUrl.RawQuery = queryParams.Encode()

	req, _ := http.NewRequest(http.MethodPost, tokenUrl.String(), nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return strava_response.Auth{}, err
	}

	if resp.StatusCode != 200 {
		return strava_response.Auth{}, errors.New("strava token request failed")
	}

	var auth strava_response.Auth
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &auth)

	if err != nil {
		return strava_response.Auth{}, err
	}

	return auth, nil
}
