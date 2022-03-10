package wallet

import "time"

type User struct {
	Reference       string    `json:"reference,omitempty" bson:"reference"`
	FirstName       string    `json:"first_name,omitempty" bson:"first_name""`
	LastName        string    `json:"last_name,omitempty" bson:"last_name""`
	Email           string    `json:"email,omitempty" bson:"email"`
	BVN             string    `json:"-" bson:"bvn"`
	Currency        string    `json:"currency"bson:"currency"`
	Password        string    `json:"-"bson:"password"`
	HashedSecretKey string    `json:"-"bson:"hashed_secret_key"`
	DateOfBirth     string    `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}
