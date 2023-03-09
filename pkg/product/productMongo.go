package product

import (
	"context"
	"fmt"
	"gorillamux/pkg/common/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

/*
BSON : Binary Encoded Json. Includes sdditional types e.g. int, long, date, floating point
bson.D : ordered document bson.D{{"hello", "world"},{"foo","bar"}}
bson.M : unordered document/map bson.M{"name":"Kyle"}
bson.A : array bson.A{"Peter","Michael"}
bson.E : ususlly used as an element inside bson.D
*/

var client *mongo.Client

// var err error
var ctx = context.Background()

func DisconnectMongo(w http.ResponseWriter, r *http.Request) {

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Closed connection")
}

func InsertOneMongo(w http.ResponseWriter, r *http.Request) {
	//var data models.User
	data := models.User{
		Name:    "name1",
		Surname: "surname1",
		Age:     22,
	}
	collection := client.Database("users").Collection("user")

	insertResult, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data had been inserted: ", insertResult.InsertedID)

}

func ReadOneMongo(w http.ResponseWriter, r *http.Request) {

}
