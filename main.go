package main

import (
	"encoding/json"
	"log"
	"net/http"
	"notes-app/database"
	"notes-app/models"

	"github.com/julienschmidt/httprouter"
)

// var collection *mongo.Collection

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.InsertUser(user)

}

func main() {
	router := httprouter.New()
	API_BASE_URL := "/go-api"
	database.CreateDBInstance()
	router.POST(API_BASE_URL+"/", Index)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
