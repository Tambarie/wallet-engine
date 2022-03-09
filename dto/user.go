package dto

type User struct {
	ID              string `json:"id,omitempty" bson:"id"`
	FirstName       string `json:"first_name,omitempty" bson:"first_name""`
	LastName        string `json:"last_name,omitempty" bson:"last_name""`
	Email           string `json:"email,omitempty" bson:"email"`
	BVN             string `json:"bvn" bson:"bvn"`
	Currency        string `json:"currency"bson:"currency"`
	SecretKey       int64  `json:"-" bson:"secret_key"`
	HashedSecretKey string `bson:"hashed_secret_key"`
	DateOfBirth     int64  `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
}
