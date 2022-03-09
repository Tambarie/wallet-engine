package wallet

import (
	"encoding/json"
	"github.com/Tambarie/wallet-engine/dto"
)

type User struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	BVN             string `json:"bvn"`
	Currency        string `json:"currency"`
	SecretKey       int64  `json:"-"`
	HashedSecretKey string `bson:"hashed_secret_key"`
	DateOfBirth     int64  `json:"date_of_birth"`
	CreatedAt       string `json:"created_at"`
}

func (u *User) ToDtoResponse() dto.User {
	return dto.User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		BVN:         u.BVN,
		Currency:    u.Currency,
		SecretKey:   u.SecretKey,
		DateOfBirth: u.DateOfBirth,
		CreatedAt:   u.CreatedAt,
	}
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		BVN         string `json:"bvn"`
		Currency    string `json:"currency"`
		SecretKey   int64  `json:"-"`
		DateOfBirth int64  `json:"date_of_birth"`
		CreatedAt   string `json:"created_at"`
	}{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		BVN:         u.BVN,
		Currency:    u.Currency,
		SecretKey:   u.SecretKey,
		DateOfBirth: u.DateOfBirth,
		CreatedAt:   u.CreatedAt,
	})
}

type Repository interface {
	CreateWallet(u *User) (*User, error)
}
