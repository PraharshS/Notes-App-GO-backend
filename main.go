package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes-app/database"
	"notes-app/models"

	"github.com/gorilla/mux"
)

// var collection *mongo.Collection

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("create hit with user", user)
	// this function is called two times on its own , so returning out from the first request having empty data
	if user.Username == "" {
		return
	}
	database.InsertUser(user)
	fmt.Println("user1", user)
	json.NewEncoder(w).Encode(user)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("login hit with user", user)

	var loggedUser = database.CheckUserLogin(user)
	json.NewEncoder(w).Encode(loggedUser)
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("create task hit with task ", task)
	if task.Message == "" {
		return
	}
	database.InsertTask(task)
	json.NewEncoder(w).Encode(task)
}
func FetchTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("fetch tasks hit")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("fetch tasks hit with user", user)
	var tasksList = database.GetTasksByUser(user)
	json.NewEncoder(w).Encode(tasksList)
}
func main() {
	router := mux.NewRouter()
	API_BASE_URL := "/go-api"
	database.CreateDBInstance()
	router.HandleFunc(API_BASE_URL+"/user/add", CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/user/login", LoginUser).Methods("POST", "OPTIONS")

	router.HandleFunc(API_BASE_URL+"/note", CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/notes-by-user", FetchTasks).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe("127.0.0.1:8002", router))
}
