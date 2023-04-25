package oauth

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	stravaUrl, _ := url.Parse("https://www.strava.com/api/v3/oauth/authorize?response_type=code&scope=activity%3Aread_all")

	query := stravaUrl.Query()
	query.Add("client_id", os.Getenv("STRAVA_CLIENT_ID"))
	query.Add("redirect_uri", os.Getenv("API_URL")+"/oauth-redirect")
	stravaUrl.RawQuery = query.Encode()

	c.Redirect(http.StatusTemporaryRedirect, stravaUrl.String())
}
