package database

import (
	"context"
	"fmt"
	"log"
	"notes-app/models"

	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println("collection instance created")

}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func InsertUser(user models.User) {
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
	fmt.Println("USER CREATED", insertResult.InsertedID)
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
func InsertTask(task models.Task) {
	collection = mongoClient.Database("golang-db").Collection("notes")
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TASK CREATED", insertResult.InsertedID)
}
func GetTasksByUser(user models.User) []models.Task {
	collection = mongoClient.Database("golang-db").Collection("notes")
	var tasksList []models.Task
	findResult, err := collection.Find(context.TODO(), bson.D{{"user.username", user.Username}})
	if err != nil {
		log.Fatal(err)
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
