package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// ConnectRedis establishes a connection to the Redis database at the specified address
// without a password and using the default database.
// It returns a pointer to the Redis client.
func ConnectRedis() *redis.Client {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})


	redisClient.Ping(ctx)

	log.Print("Connected to Redis")
	return redisClient
}
