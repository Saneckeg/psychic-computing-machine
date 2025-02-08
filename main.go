package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type MessageResponse struct {
	Task string `json:"task"`
	Id   bool   `json:"id"`
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages{id}", PatchMessages).Methods("PATCH")
	router.HandleFunc("/api/messages{id}", DeleteMessages).Methods("DELETE")
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

func PatchMessages(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var message Message
	if err := DB.First(&message, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := DB.Model(&message).Updates(updates)
	if result.Error != nil {
		http.Error(w, "Ошибка обновления в БД", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(message)
}

func DeleteMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var message Message
	if err := DB.First(&message, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if err := DB.Delete(&message).Error; err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Message deleted")
}
