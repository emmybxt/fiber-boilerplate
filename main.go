package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"waitlist/db"
	"waitlist/middleware"
	"waitlist/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	port := getPort()

	database := db.ConnectDatabase()
	redis := db.ConnectRedis()

	requestLogger := &middleware.RLogger{}
	router.Use(requestLogger.LogRequest)

	// Setup routes
	routes.SetupRoutes(router, database, redis)

	listenAddress := fmt.Sprintf(":%s", port)
	err = router.Run(listenAddress)
	if err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	return port
}
