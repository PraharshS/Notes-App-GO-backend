package database

import (
	"context"
	"fmt"
	"log"
	"notes-app/models"
	encryption "notes-app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client
var DBName string
var userCollectionName string
var taskCollectionName string
var userCollection *mongo.Collection
var taskCollection *mongo.Collection

func InsertUser(user models.User) models.User {
	HashedUserPassword, err := encryption.HashPassword(user.Password)
	user.Password = HashedUserPassword

	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("USER CREATED ID", insertResult.InsertedID, user.Username)
	var result models.User
	var nullUser models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	fmt.Println("RESULT", result.ID, result.Username)
	fmt.Println("RESULT", result)
	if err != nil {
		return nullUser
	}
	fmt.Println("ERR")
	return result
}
func CheckUserLogin(user models.User) models.User {
	var result models.User
	var nullUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
	if err != nil {
		return nullUser
	}
	var passwordMatch = encryption.CheckPasswordHash(user.Password, result.Password)
	if !passwordMatch {
		return nullUser
	}
	fmt.Println("Matched Data", result.ID, result.Username, result.Password)
	return result
}
func InsertTask(task models.Task) models.Task {
	taskCollection = mongoClient.Database(DBName).Collection("tasks")
	insertResult, err := taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	var result models.Task
	err = taskCollection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TASK CREATED", insertResult.InsertedID)
	return result
}
func GetTasksByUser(userIDHex string) []models.Task {
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		panic(err)
	}
	var tasksList []models.Task
	findResult, err := taskCollection.Find(context.TODO(), bson.M{"user._id": userID})
	if err != nil {
		log.Fatal(err)
		return tasksList
	}
	for findResult.Next(context.TODO()) {
		//Create a value intwhich the single document can be decoded
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

	deleteResult, _ := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": taskId})
	if deleteResult.DeletedCount == 0 {
		log.Fatal("Error on deleting one Task", err)

	}
	fmt.Println("Deleted task of Id ", taskIdHex, deleteResult)
}
func ToggleTaskDone(taskIdHex string) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)

	var foundTask models.Task
	err = taskCollection.FindOne(context.TODO(), bson.M{"_id": taskId}).Decode(&foundTask)
	var toggleStatusTask = !foundTask.IsDone
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(foundTask)
	result, _ := taskCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": taskId},
		bson.M{"$set": bson.M{"is_done": toggleStatusTask}},
	)
	fmt.Println("Task done with id", taskIdHex, result)
}
func UpdateTask(taskIdHex string, updatedTask models.Task) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	filter := bson.M{"_id": taskId}
	result, err := taskCollection.ReplaceOne(context.TODO(), filter, updatedTask)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"insert: %d, updated%d, deleted: %d /n",
		result.MatchedCount,
		result.ModifiedCount,
		result.UpsertedCount,
	)
}
