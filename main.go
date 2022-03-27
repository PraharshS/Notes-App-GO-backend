package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-app/db"
	"notes-app/router"
)

func main() {
	db.CreateDBInstance()
	r := router.Router()
	fmt.Println("starting the server on port 8000...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}
