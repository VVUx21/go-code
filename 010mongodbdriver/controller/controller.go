package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/VVUx21/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
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

//get all records from mongodb
func GetAllRecords() ([]primitive.M, error) {
	var movies []primitive.M
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())	
	return movies, err
}

//ACtual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies, err := GetAllRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertResult, err := InsertOneRecord(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(insertResult)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	movieID := mux.Vars(r)["id"]
	updateResult, err := UpdateOneRecord(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updateResult)
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	movieID := mux.Vars(r)["id"]
	deleteResult, err := DeleteOneRecord(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	deleteResult, err := DeleteAllRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}