package router

import (
	service "notes-app/service"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	API_BASE_URL := "/go-api"
	router := mux.NewRouter()
	router.HandleFunc(API_BASE_URL+"/user/add", service.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/user/login", service.LoginUser).Methods("POST", "OPTIONS")

	router.HandleFunc(API_BASE_URL+"/task", service.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/tasksByUser/{userID}", service.FetchTasks).Methods("GET", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/{id}", service.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/{id}", service.ToggleTaskDone).Methods("PUT", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/update-task/{id}", service.UpdateTask).Methods("PUT", "OPTIONS")
	return router
}
