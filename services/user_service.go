package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"waitlist/middleware"
	"waitlist/models"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct {
	db *mongo.Database
	log *slog.Logger
	redis *redis.Client
}


func NewUserService(db *mongo.Database, redis *redis.Client) *UserService {
	return &UserService{
		db: db,
		log: slog.Default(),
		redis: redis,
	}
}

func (us *UserService) GetAllUsers() ([]models.UserModel, error) {
	var users []models.UserModel

	collection := us.db.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	us.log.Info("Users are returned")
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) AddUser(user *models.UserModel) (*models.UserModel, error) {
	
	usersCollection := us.db.Collection("users")

	// Check if the email already exists
	emailFilter := bson.M{"email": user.Email}
	emailCount, err := usersCollection.CountDocuments(context.Background(), emailFilter)
	if err != nil {
		return nil, err
	}
	if emailCount > 0 {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user object with hashed password and generated ID
	newUser := &models.UserModel{
		Id:        primitive.NewObjectID(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	
	_, err = usersCollection.InsertOne(context.Background(), newUser)
	
	if err != nil {
		return nil, err
	}

	return newUser, nil
}


func (us *UserService) Login(user *models.UserModel) (*models.UserModel, string, error) {
	usersCollection := us.db.Collection("users")

	ctx := context.Background()
	userFilter := bson.M{"email": user.Email}
	var dbUser models.UserModel
	err := usersCollection.FindOne(ctx, userFilter).Decode(&dbUser)
	fmt.Print(err)
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, "", err
	}

	token, err := middleware.GenerateJWT(&dbUser)
	if err != nil {
		return nil, "", err
	}

	redisKey := fmt.Sprintf("user:%s", dbUser.Email)
	err = us.redis.Set(ctx, redisKey, token, 0).Err()
	if err != nil {
		return nil, "", err
	}

	return &dbUser, token, nil
}
