package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// ConnectDatabase establishes a connection to the MongoDB database specified by the environment variables DB_URI and DB_NAME.
// It returns a pointer to the connected database.
func ConnectDatabase() *mongo.Database {
	dbURI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Connected to database")
	database := client.Database(dbName)
	return database
}
