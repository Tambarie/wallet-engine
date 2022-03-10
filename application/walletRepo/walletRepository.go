package walletRepo

import (
	"fmt"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
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

func (walletRepo *RepositoryDB) GetUserByEmail(email string) ([]*wallet.User, error) {
	coll := walletRepo.db.Database("opay").Collection("opay-collection")
	filter := bson.D{{"email", email}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*wallet.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (walletRepo *RepositoryDB) CreateWallet(u *wallet.User) (*wallet.User, error) {
	coll := walletRepo.db.Database("opay").Collection("opay-collection")
	result, err := coll.InsertOne(context.TODO(), u)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return u, err
}

func (walletRepo *RepositoryDB) CheckIfPasswordExists(userReference string) ([]*wallet.User, error) {
	coll := walletRepo.db.Database("opay").Collection("opay-collection")
	filter := bson.D{{"reference", userReference}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*wallet.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (walletRepo *RepositoryDB) SaveTransaction(t *wallet.Transaction) (*wallet.Transaction, error) {
	coll := walletRepo.db.Database("opay").Collection("transaction-collection")
	result, err := coll.InsertOne(context.TODO(), t)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return t, err
}

func (walletRepo *RepositoryDB) PostToAccount(a *wallet.Wallet) (*wallet.Wallet, error) {
	coll := walletRepo.db.Database("opay").Collection("account-balance")
	result, err := coll.InsertOne(context.TODO(), a)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return a, err
}

func (walletRepo *RepositoryDB) GetAccountBalance(userID string) ([]*wallet.Wallet, error) {
	coll := walletRepo.db.Database("opay").Collection("account-balance")

	filter := bson.D{{"userid", userID}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*wallet.Wallet
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err

}

func (walletRepo *RepositoryDB) ChangeUserStatus(isActive bool, userReference string) (interface{}, error) {
	coll := walletRepo.db.Database("opay").Collection("opay-collection")

	filter := bson.D{{"reference", userReference}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"is_active", isActive}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")

	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	return result.UpsertedID, err
}
