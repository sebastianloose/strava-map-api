package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/sebastianloose/strava-map-api/controller"
	"github.com/sebastianloose/strava-map-api/controller/oauth"
	"github.com/sebastianloose/strava-map-api/service/cache"

	"github.com/gin-gonic/gin"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	go cache.StartUserCacheWorker()
	go cache.StartActivityCacheWorker()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Origin", "Accept", "authorization", "X-Requested-With", "Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/oauth-redirect", oauth.Redirect)
	router.GET("/oauth-login", oauth.Login)
	router.GET("/activities", controller.GetActivities)
	router.GET("/activity/:activityId", controller.GetActivity)
	router.Run()

	fmt.Println("Server listening on port 8080")
}
