package mongoDB

import (
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

func Init() *mongo.Client {

	mongoURL := fmt.Sprintf("%s://%s:%s", os.Getenv("db_type"), os.Getenv("mongo_db_host"), os.Getenv("mongo_db_port"))

	mongoTimeout := time.Minute * 15

	// using go mongo-driver  to connect to mongoDB

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout))
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("error %v", err)
	}

	log.Println("Database Connected Successfully...")

	return client
}
