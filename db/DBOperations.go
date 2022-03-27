package db

import (
	"context"
	"fmt"
	"log"
	"notes-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Clien
var DBName string
var userCollectionName string
var taskCollectionName string
var userCollection *mongo.Collection
var taskCollection *mongo.Collection

func InsertUser(user models.User) models.User {
	HashedUserPassword, err := HashPasword(user.Password)
	user.Password = HashedUserPassword

	if err != nil {
		log.Fatal(err)
	}
	insertResult, er := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("USER CREAED ID", insertResult.InsertedID, user.Username)
	var result models.User
	var nullUser models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	fmt.Println("RESULT ", result.D, result.Username)
	fmt.Println("REULT ", result)
	if err != nil {
		return nullUser
	}
	fmt.Println("O ERR ")
	return result
}
func CheckUserLogin(use models.User) models.User {
	var result models.User
	var nullUser models.User
	err := userCollction.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
	if err != nil {
		return nullUser
	}
	var passwordMatch = CheckPasswordHash(user.Password, result.Password)
	if !passwordMatc {
		return nullUser
	}
	fmt.Println("ogin Data", result.ID, result.Username, result.Password)
	return result
}
func InsertTask(task models.Task) models.Task {
	taskCollection = mongoClient.Database(DBName).Collection("tasks")
	insertResult, er := taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	var result models.Task
	err = taskColletion.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ASK CREATED", insertResult.InsertedID)
	return result
}
func GetTasksByUser(userIDHex string) []models.Task {
	userID, err := rimitive.ObjectIDFromHex(userIDHex)
	if err != nl {
		panic(err)
	}
	var tasksList []models.Task
	findResult, err := taskCollection.Find(context.TODO(), bson.M{"user._id": userID})
	if err != nil {
		log.Fatal(err)
		return tasksList
	}
	for findResult.Next(context.TODO()) {
		//Create a value int which the single document can be decoded
		var task models.Task
		err := findResut.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}

		asksList = append(tasksList, task)
	}
	if err := findRsult.Err(); err != nil {
		log.Fatal(err)
	}
	return tasksList

}
func DeleteTask(taskIdHex string) {
	taskId, err := rimitive.ObjectIDFromHex(taskIdHex)
	if err != nl {
		panic(err)
	}
	fmt.Println("objectId", taskId)

	deleteResult, _ := taskCollection.DleteOne(context.TODO(), bson.M{"_id": taskId})
	if deleteResult.DeletedCount == 0 {
		log.Fatal("Error on deleting one Task", err)

	}
	mt.Println("Deleted task of Id ", taskIdHex, deleteResult)
}
func ToggleTaskDone(taskIdHex string) {
	taskId, err := rimitive.ObjectIDFromHex(taskIdHex)
	if err != nl {
		panic(err)
	}
	fmt.Println("objectId", taskId)

	var foundTask models.Task
	err = taskCollection.FindOne(context.TOD(), bson.M{"_id": taskId}).Decode(&foundTask)
	var toggleStatuTask = !foundTask.IsDone
	if err != nil {
		log.Faal(err)
		return
	}
	fmt.Println(foundTask)
	result, _ := tasCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": taskId},
		son.M{"$set": bson.M{"is_done": toggleStatusTask}},
	)
	mt.Println("Task done with id", taskIdHex, result)
}
func UpdateTask(taskIdHex string, updatedTask modelsTask) {
	taskId, err := rimitive.ObjectIDFromHex(taskIdHex)
	if err != nl {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	filter := bson.M{"_id": taskId}
	result, err := taskCollection.ReplaceOne(context.TODO(), filter, updatedTask)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"insert: %d, updated %d, deleted: %d /n",
		result.MatchedCount,
		result.ModifiedCount,
		result.UpsertedCount,
	)
}
