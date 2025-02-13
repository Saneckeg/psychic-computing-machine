package main

import (
	"SecondProject/internal/database"
	"SecondProject/internal/handlers"
	"SecondProject/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type MessageResponse struct {
	Task string `json:"task"`
	Id   bool   `json:"id"`
}

func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/messages/{id}", handler.PatchMessages).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", handler.DeleteMessages).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
