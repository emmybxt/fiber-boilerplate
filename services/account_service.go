package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"waitlist/models"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	db    *mongo.Database
	log   *slog.Logger
	redis *redis.Client
}

func NewAccountService(db *mongo.Database, redis *redis.Client) *AccountService {
	return &AccountService{
		db:    db,
		log:   slog.Default(),
		redis: redis,
	}
}

func (s *AccountService) CreateAccount(user *models.UserModel) (models.Account, error) {
	collection := s.db.Collection("accounts")

	account := models.Account{
		Id:                  primitive.NewObjectID(),
		AccountName:         fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		AccountOwnerId:      user.Id,
		Currency:            "NGN",
		Iso2:                "NG",
		Balance:             100000,
		BalanceBeforeCredit: 0,
		BalanceBeforeDebit:  0,
		BalanceAfterDebit:   0,
		BalanceAfterCredit:  0,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	_, err := collection.InsertOne(context.Background(), account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func (s *AccountService) GetAccountByOwnerId(id primitive.ObjectID) (models.Account, error) {

	collection := s.db.Collection("accounts")
	filter := bson.M{"accountOwnerId": id}
	var account models.Account
	err := collection.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}
