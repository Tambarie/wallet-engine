package wallet

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryDB struct {
	db *mongo.Client
}

// NewWalletRepositoryDB function to initialize RepositoryDB struct
func NewWalletRepositoryDB(client *mongo.Client) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

func (walletRepo *RepositoryDB) Get(id string) (*User, error) {
	return nil, nil
}

func (walletRepo *RepositoryDB) Create(u *User) (*User, error) {
	return u, nil
}

func (walletRepo *RepositoryDB) GetUserByEmail(email string) (*User, error) {
	return nil, nil
}
