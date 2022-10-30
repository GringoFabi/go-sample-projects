package connector

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
    Name string
    Age  int
    City string
}

const uri = "mongodb://user:pass@localhost:27017/"

func Connect() (*mongo.Client, error) {

	fmt.Println("Init connection...")

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, err
}

func SetupData(client *mongo.Client, ) {
	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	trainers := []interface{}{ash, misty, brock}

	insertResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}

func GetTrainer(client *mongo.Client, name string) (Trainer, error) {
	collection := client.Database("test").Collection("trainers")

	filter := bson.D{{Key: "name", Value: name}}
	var result Trainer

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}