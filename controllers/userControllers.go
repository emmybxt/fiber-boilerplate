package controllers

import (
	"net/http"
	"waitlist/models"
	"waitlist/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserImpl struct {
	userService *services.UserService
	accountService *services.AccountService
}

func NewUser(db *mongo.Database, redis *redis.Client) *UserImpl {
	return &UserImpl{
		userService: services.NewUserService(db, redis),
		accountService: services.NewAccountService(db, redis),
	}
}

func (um *UserImpl) GetAllUsers(c *gin.Context) {
	users, err := um.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserImpl) AddUser(c *gin.Context) {
	var newUser models.UserModel

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	addedUser, err := u.userService.AddUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}


	_, err = u.accountService.CreateAccount(addedUser)
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An Error occurred while creating the user account"})
    }

	c.JSON(http.StatusOK, addedUser)
}

func (u *UserImpl) Login(c *gin.Context) {
	var userRequest models.UserModel
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	loggedInUser, token, err := u.userService.Login(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	userAccount, err := u.accountService.GetAccountByOwnerId(loggedInUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user account"})
		return
	}
	response := gin.H{
		"user": loggedInUser,
		"token": token,
		"account": userAccount,
	}

	c.JSON(http.StatusOK, response)
}
