package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getConnection get a connection to the PostgresqlDB
func GetConnection() *sql.DB {
	dsn := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func MongoEndPoint(collection string) *mongo.Collection {
	dsn := os.Getenv("Mongo_URL")
	database := os.Getenv("mongodb")

	clientOpts := options.Client().ApplyURI(dsn)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Check the connections
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return client.Database(database).Collection(collection)

}
