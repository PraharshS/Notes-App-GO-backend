package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes-app/db"
	"notes-app/models"

	"github.com/gorilla/mux"
)

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
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("create task hit with task ", task)
	if task.Name == "" {
		return
	}
	var createdTask = db.InsertTask(task)
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
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`userID := `, userID)
	fmt.Println("fetch tasks hit with userID", userID)
	var tasksList = db.GetTasksByUser(userID)
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
	db.DeleteTask(id)
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
	db.ToggleTaskDone(id)
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
	db.UpdateTask(id, task)
}
