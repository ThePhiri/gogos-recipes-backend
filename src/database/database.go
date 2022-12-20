package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_DB")),
	}

}

//getbyid from mongo
func GetById(collection string, id string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := MI.DB.Collection(collection)
	log.Printf("Getting %s with id %s", collection, id)

	recipeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting id to hex: %v", err)
		return nil, err
	}

	var result bson.M
	//find one from object id mongo
	err = c.FindOne(ctx, bson.M{"_id": recipeId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	log.Printf("Found document: %v", result)

	return result, nil
}

//getbyfilter from mongodb
func GetByFilter(collection string, filter bson.M) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := MI.DB.Collection(collection)
	log.Printf("Getting %s with filter %s", collection, filter)
	var results []interface{}
	//find one from object id mongo
	cur, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
