package middleware

import (
	"context"
	"final/cmd/models"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection object/instance
var lists *mongo.Collection
var tasks *mongo.Collection

// create connection with mongo db
func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	list_collection := os.Getenv("DB_LIST_COLLECTION_NAME")
	task_collection := os.Getenv("DB_TASK_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
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

	lists = client.Database(dbName).Collection(list_collection)
	tasks = client.Database(dbName).Collection(task_collection)

	fmt.Println("Collection instance created!")
}

// get all task from the DB and return it
func GetAllLists() []primitive.M {
	cur, err := lists.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// get all task from the DB and return it
func GetAllTasks(param string) []primitive.M {
	id, _ := strconv.Atoi(param)
	filter := bson.D{{Key: "listId", Value: id}}
	cur, err := tasks.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)

		if e != nil {
			log.Fatal(e)
		}
		log.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

//Insert one list in the DB
func InsertList(list models.List) {
	if list.ID == 0 {
		id := rand.Int63n(1000000)
		list.ID = id
	}

	insertResult, err := lists.InsertOne(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

//Insert one task in the DB
func InsertTask(task models.Task) {
	if task.ID == 0 {
		id := rand.Intn(1000000)
		task.ID = int64(id)
	}

	log.Println(task.ID)

	insertResult, err := tasks.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func UpdateTask(id int64) {
	fmt.Println(id)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}
	result, err := tasks.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one list from the DB, delete by ID
func DeleteOneList(list string) error {
	fmt.Println(list)
	id, _ := strconv.Atoi(list)
	filter := bson.M{"_id": id}
	d, err := lists.DeleteOne(context.Background(), filter)
	tasks.DeleteMany(context.Background(), bson.M{"listId": id})

	fmt.Println("Deleted Document", d.DeletedCount)
	if err != nil {
		return err
	}
	return nil
}

// delete one task from the DB, delete by ID
func DeleteOneTask(task string) error {
	fmt.Println(task)
	id, _ := strconv.Atoi(task)
	filter := bson.M{"_id": id}
	d, err := tasks.DeleteOne(context.Background(), filter)

	fmt.Println("Deleted Document", d.DeletedCount)
	if err != nil {
		return err
	}
	return nil
}
