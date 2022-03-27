package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDBInstance() {
	godotenv.Load()
	connectionString := os.Getenv("DB_URI")
	DBName = os.Getenv("DB_NAME")
	userCollectionName = os.Getenv("USER_DBCOLLECTION_NAME")
	taskCollectionName = os.Getenv("TASK_DBCOLLECTION_NAME")

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
	userCollection = client.Database(DBName).Collection(userCollectionName)
	taskCollection = client.Database(DBName).Collection(taskCollectionName)
}
