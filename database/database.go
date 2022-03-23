package database

import (
	"context"
	"fmt"
	"log"
	"notes-app/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func InsertUser(user models.User) {

	insertResult, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("USER CREATED", insertResult.InsertedID)
}
