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
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://zitrakz:mongopass747@pdgcluster.txsie.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	// client, err = mongo.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
	// fmt.Println("connected to mongo DB")
	// collection := client.Database("golang-db").Collection("users")

}
func InsertUser(user models.User) {

	insertResult, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single record", insertResult.InsertedID)
}
