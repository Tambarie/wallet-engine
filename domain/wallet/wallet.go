package wallet

import (
	"encoding/json"
	"github.com/Tambarie/wallet-engine/dto"
)

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password" sql:"-"`
	HashPassword string `json:"-" sql:"password"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type Repository interface {
	Get(id string) (*User, error)
	Create(u *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

func (u *User) ToDtoResponse() dto.User {
	return dto.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID           string `json:"id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		HashPassword string `json:"-"`
		CreatedAt    int64  `json:"created_at"`
		UpdatedAt    int64  `json:"updated_at"`
	}{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	})
}
