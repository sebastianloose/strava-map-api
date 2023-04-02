package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sebastianloose/strava-map-api/controller"
	"github.com/sebastianloose/strava-map-api/controller/oauth"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/oauth-redirect", oauth.Redirect)
	router.GET("/oauth-login", oauth.Login)
	router.GET("/activities", controller.GetActivities)
	router.Run()

	fmt.Println("Server listening on port 8080")
}
