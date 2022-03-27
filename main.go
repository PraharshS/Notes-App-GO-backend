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
	var addedUser = database.InsertUser(user)
	// fmt.Println("added user ", addedUser)
	json.NewEncoder(w).Encode(addedUser)
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

	// Preflight request sent by react
	if r.Method == "OPTIONS" {
		return
	}
	var task models.Task
	fmt.Println("create task hit with task ", task.Targeted_Date)
	json.NewDecoder(r.Body).Decode(&task)
	// DATE,_ := time.Parse("YYYY-MM-DD", task.Targeted_Date)
	// fmt.Println("create task hit with task date ", DATE)
	if task.Message == "" {
		return
	}
	var createdTask = database.InsertTask(task)
	json.NewEncoder(w).Encode(createdTask)
}
func FetchTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Preflight request sent by react
	if r.Method == "OPTIONS" {
		return
	}
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("fetch tasks hit with user", user)
	var tasksList = database.GetTasksByUser(user)
	if tasksList == nil {
		tasksList = make([]models.Task, 0)
	}
	json.NewEncoder(w).Encode(tasksList)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Preflight request sent by react
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println("delete hit")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	database.DeleteTask(id)
	json.NewEncoder(w).Encode("Task Deleted Successfully")
}
func ToggleTaskDone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Done task hit")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	database.ToggleTaskDone(id)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Preflight request sent by react
	if r.Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`update id := `, id)
	var task models.Task
	fmt.Println("update task hit with task ", task)
	json.NewDecoder(r.Body).Decode(&task)
	database.UpdateTask(id, task)
}
func main() {
	router := mux.NewRouter()
	API_BASE_URL := "/go-api"
	database.CreateDBInstance()
	router.HandleFunc(API_BASE_URL+"/user/add", CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/user/login", LoginUser).Methods("POST", "OPTIONS")

	router.HandleFunc(API_BASE_URL+"/task", CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/tasksByUser", FetchTasks).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/{id}", DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/{id}", ToggleTaskDone).Methods("PUT", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/update-task/{id}", UpdateTask).Methods("PUT", "OPTIONS")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
