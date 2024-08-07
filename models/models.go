package models

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WaitlistEntry struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"email" json:"email"required:"true"`
}

type UserModel struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email" json:"email" required:"true"`
	FirstName string             `bson:"firstName" json:"firstname" required:"true"`
	LastName  string             `bson:"lastName" json:"lastname" required:"true"`
	Password  string             `bson:"password" json:"password" required:"true"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" default:"now"`
}

func (u UserModel) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

type TransactionType string
type TransactionStatus string

const (
	Pending   TransactionStatus = "PENDING"
	Completed TransactionStatus = "completed"
	Failed    TransactionStatus = "failed"
	Reversed  TransactionStatus = "reversed"
)

const (
	TransactionTypeCredit TransactionType = "CREDIT"
	TransactionTypeDebit  TransactionType = "DEBIT"
)

type Account struct {
	CreatedAt           time.Time          `bson:"createdAt"`
	UpdatedAt           time.Time          `bson:"updatedAt"`
	Id                  primitive.ObjectID `bson:"_id"`
	AccountName         string             `bson:"accountName"`
	AccountOwnerId      primitive.ObjectID `bson:"accountOwnerId"`
	Currency            string             `bson:"currency"`
	Iso2                string             `bson:"iso2"`
	Balance             float64            `bson:"balance"`
	BalanceBeforeCredit float64            `bson:"balanceBeforeCredit"`
	BalanceBeforeDebit  float64            `bson:"balanceBeforeDebit"`
	BalanceAfterDebit   float64            `bson:"balanceAfterDebit"`
	BalanceAfterCredit  float64            `bson:"balanceAfterCredit"`
}

type Transaction struct {
	TimeCreated            time.Time          `bson:"timeCreated"`
	TimeUpdated            time.Time          `bson:"timeUpdated"`
	DestinationAccountId   string             `bson:"destinationAccountId"`
	DestinationAccountName string             `bson:"destinationAccountName"`
	Type                   TransactionType    `bson:"type"`
	Status                 TransactionStatus  `bson:"status"`
	SourceAccountId        string             `bson:"sourceAccountId"`
	SourceAccountName      string             `bson:"sourceAccountName"`
	Id                     primitive.ObjectID `bson:"_id"`
	Reference              string             `bson:"reference"`
	Description            string             `bson:"description"`
	BalanceBeforeDebit     float64            `bson:"balanceBeforeDebit"`
	BalanceAfterDebit      float64            `bson:"balanceAfterDebit"`
	BalanceAfterCredit     float64            `bson:"balanceAfterCredit"`
	BalanceBeforeCredit    float64            `bson:"balanceBeforeCredit"`
	Amount                 float64            `bson:"amount"`
}

type TransactionPayload struct {
	FromAccountId string  `json:"fromAccountId"`
	ToAccountId   string  `json:"toAccountId"`
	Description   string  `json:"description"`
	Amount        float64 `json:"amount"`
}

func (t TransactionPayload) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.FromAccountId, validation.Required),
		validation.Field(&t.ToAccountId, validation.Required),
		validation.Field(&t.Amount, validation.Required),
	)
}

func (t TransactionType) IsValid() bool {
	switch t {
	case TransactionTypeCredit, TransactionTypeDebit:
		return true
	default:
		return false
	}
}

func (t TransactionType) String() string {
	return string(t)
}

func (t TransactionStatus) IsValid() bool {
	switch t {
	case Pending, Completed, Failed, Reversed:
		return true
	default:
		return false
	}
}

func (t TransactionStatus) String() string {
	return string(t)
}
