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
	database.InsertUser(user)
}
func doNothing(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	API_BASE_URL := "/go-api"
	database.CreateDBInstance()
	router.HandleFunc("/favicon.ico", doNothing)
	router.HandleFunc(API_BASE_URL+"/user/add", CreateUser).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
