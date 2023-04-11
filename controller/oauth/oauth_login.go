package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	clientId := url.QueryEscape(os.Getenv("STRAVA_CLIENT_ID"))
	redirectUrl := url.QueryEscape(os.Getenv("API_URL") + "/oauth-redirect")

	stravaUrl := fmt.Sprintf("https://www.strava.com/api/v3/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=activity%%3Aread_all", clientId, redirectUrl)
	c.Redirect(http.StatusTemporaryRedirect, stravaUrl)
}
