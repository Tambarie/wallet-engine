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

func (walletRepo *RepositoryDB) CreateWallet(u *User) (*User, error) {
	return nil, nil
}
