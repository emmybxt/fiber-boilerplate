package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigMongoDB() *config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &config{
		host:     os.Getenv("MONGODB_HOST"),
		database: os.Getenv("MONGODB_HOST"),
		// password:   os.Getenv("MONGODB_ROOT_PASSWORD"),
		// user:       os.Getenv("MONGODB_ROOT_USER"),
		ctxTimeout: 60 * time.Second,
	}
}

func newConfigPostgres() *config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	fmt.Print(os.Getenv("POSTGRES_HOST"))
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
