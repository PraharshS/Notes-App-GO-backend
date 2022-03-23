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

func CreateDBInstance() {
	connectionString := `mongodb+srv://zitrakz:mongopass747@pdgcluster.txsie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority`
	dbName := `golang-db`
	collName := `users`

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

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
func CheckUserLogin(user models.User) {
	var result models.User
	err := collection.FindOne(context.Background(), bson.D{{"username", user.Username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Fatal(err)
		return
	}
	// CheckPasswordHash(user.Password, result[fmt.Sprint("password")])
	fmt.Println("found document ", result)
	fmt.Println("found values " + result.Username + " " + result.Password)
}
