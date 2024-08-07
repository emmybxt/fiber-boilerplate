package routes

import (
	"waitlist/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// type routes struct {
// 	user     *controllers.UserImpl
// 	waitlist *controllers.Waitlist
// 	db *mongo.Database
// }

// func NewRoutes(user *controllers.UserImpl, waitlist *controllers.Waitlist, db *mongo.Database) *routes {
// 	return &routes{
// 		user:     user,
// 		waitlist: waitlist,
// 		db: db,
// 	}
// }

// SetupRoutes configures the routes for the API.
func SetupRoutes(router *gin.Engine, db *mongo.Database, redis *redis.Client) {

	userController := controllers.NewUser(db, redis)
	waitlistController := controllers.NewWaitlist(db, redis)

	router.POST("/api/waitlist", waitlistController.AddToWaitlist())
	router.GET("/api/waitlist", waitlistController.GetWaitList())
	router.DELETE("/api/waitlist/:email", waitlistController.DeleteFromWaitlist())

	userRoutes := router.Group("/api/users")
	userRoutes.GET("/", userController.GetAllUsers)
	userRoutes.POST("/", userController.AddUser)
	userRoutes.POST("/login", userController.Login)
}
