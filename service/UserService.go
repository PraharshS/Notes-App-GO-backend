package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes-app/database"
	"notes-app/models"
	encrytion "notes-app/util"
)

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
	json.NewEncoder(w).Encode(addedUser)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	token := encrytion.TokenGenerator()
	fmt.Println("login Token", token)
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("login hit with user", user)
	var loggedUser = database.CheckUserLogin(user)
	var userWithToken models.UserWithToken
	userWithToken.Token = token
	userWithToken.User = loggedUser

	json.NewEncoder(w).Encode(userWithToken)
}
