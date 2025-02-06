package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type MessageResponse struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
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

	// Устанавливаем заголовок Content-Type в application/json
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON в ответ
	json.NewEncoder(w).Encode(req)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message

	DB.Find(&messages)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}
