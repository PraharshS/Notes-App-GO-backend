package database

import (
	"context"
	"fmt"
	"log"
	"notes-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection

var mongoClient *mongo.Client

func CreateDBInstance() {
	connectionString := `mongodb+srv://zitrakz:mongopass747@pdgcluster.txsie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority`
	dbName := `golang-db`
	collName := `users`

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	mongoClient = client
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")
	collection = client.Database(dbName).Collection(collName)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func InsertUser(user models.User) models.User {
	collection = mongoClient.Database("golang-db").Collection("users")
	HashedUserPassword, err := HashPassword(user.Password)
	user.Password = HashedUserPassword

	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("USER CREATED ID", insertResult.InsertedID, user.Username)
	var result models.User
	var nullUser models.User
	err = collection.FindOne(context.Background(), bson.D{{"_id", insertResult.InsertedID}}).Decode(&result)
	fmt.Println("RESULT ", result.ID, result.Username)
	fmt.Println("RESULT ", result)
	if err != nil {
		return nullUser
	}
	fmt.Println("NO ERR ")
	return result
}
func CheckUserLogin(user models.User) models.User {
	collection = mongoClient.Database("golang-db").Collection("users")
	var result models.User
	var nullUser models.User
	err := collection.FindOne(context.Background(), bson.D{{"username", user.Username}}).Decode(&result)
	if err != nil {
		return nullUser
	}
	var passwordMatch = CheckPasswordHash(user.Password, result.Password)
	if !passwordMatch {
		return nullUser
	}
	fmt.Println("Login Data", result.ID, result.Username, result.Password)
	return result
}
func InsertTask(task models.Task) models.Task {
	collection = mongoClient.Database("golang-db").Collection("notes")
	insertResult, err := collection.InsertOne(context.Background(), task)
	var result models.Task
	err = collection.FindOne(context.Background(), bson.D{{"_id", insertResult.InsertedID}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TASK CREATED", insertResult.InsertedID)
	return result
}
func GetTasksByUser(user models.User) []models.Task {
	collection = mongoClient.Database("golang-db").Collection("notes")
	var tasksList []models.Task
	findResult, err := collection.Find(context.TODO(), bson.D{{"user.username", user.Username}})
	if err != nil {
		log.Fatal(err)
		return tasksList
	}
	for findResult.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var task models.Task
		err := findResult.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}

		tasksList = append(tasksList, task)

	}
	if err := findResult.Err(); err != nil {
		log.Fatal(err)
	}
	return tasksList

}
func DeleteTask(taskIdHex string) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	collection = mongoClient.Database("golang-db").Collection("notes")

	deleteResult, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": taskId})
	if deleteResult.DeletedCount == 0 {
		log.Fatal("Error on deleting one Hero", err)

	}
	fmt.Println("Deleted task of Id ", taskIdHex, deleteResult)
}
func ToggleTaskDone(taskIdHex string) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	collection = mongoClient.Database("golang-db").Collection("notes")
	var foundTask models.Task
	err = collection.FindOne(context.TODO(), bson.D{{"_id", taskId}}).Decode(&foundTask)
	var toggleStatusTask = !foundTask.IsDone
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(foundTask)
	result, _ := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": taskId},
		bson.M{"$set": bson.M{"is_done": toggleStatusTask}},
	)
	fmt.Printf("Task done with id", taskIdHex, result)
}
func UpdateTask(taskIdHex string, updatedTask models.Task) {
	collection = mongoClient.Database("golang-db").Collection("notes")
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	filter := bson.M{"_id": taskId}
	result, err := collection.ReplaceOne(context.TODO(), filter, updatedTask)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"insert: %d, updated: %d, deleted: %d /n",
		result.MatchedCount,
		result.ModifiedCount,
		result.UpsertedCount,
	)
}
