package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/VVUx21/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "mydatabase"
const collectionName = "mycollection"

// MOST IMPORtant
var collection *mongo.Collection

//connect to mongodb
func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	//explanation context.todo() is used when you are not sure which context to use or when you don't have a specific context available.
	//It returns a non-nil, empty Context. It is typically used in main functions, initialization, and tests, and as the top-level Context for incoming requests.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successful")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

//mongodb helpers - file
func InsertOneRecord(movie model.Netflix) (*mongo.InsertOneResult, error) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single record ", insertResult.InsertedID)
	return insertResult, err
}

func UpdateOneRecord(movieID string) (*mongo.UpdateResult, error) {
	id , _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated a single record ", updateResult.ModifiedCount)
	return updateResult, err
}

func DeleteOneRecord(movieID string) (*mongo.DeleteResult, error) {
	id , _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}	
	fmt.Println("Deleted a single record ", deleteResult.DeletedCount)
	return deleteResult, err
}

func DeleteAllRecords() (*mongo.DeleteResult, error) {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter,nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all records ", deleteResult.DeletedCount)
	return deleteResult, err
}