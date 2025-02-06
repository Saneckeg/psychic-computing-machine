package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"Message"`
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)

}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req Message

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := DB.Create(&req)
	if result.Error != nil {
		http.Error(w, "Ошибка записи в БД", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Задание обновлено: %s %t", req.Task, req.IsDone)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var tasks []string

	DB.Model(&Message{}).Pluck("task", &tasks)

	json.NewEncoder(w).Encode(tasks)
}
