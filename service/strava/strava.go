package strava

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

func TokenExchange(userCode string) (UserToken, error) {
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

	req, _ := http.NewRequest("POST", tokenUrl.String(), nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return UserToken{}, err
	}

	if resp.StatusCode != 200 {
		return UserToken{}, errors.New("strava token request not successful")
	}

	var userToken UserToken
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &userToken)

	if err != nil {
		return UserToken{}, err
	}

	return userToken, nil
}
