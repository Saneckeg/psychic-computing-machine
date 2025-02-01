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
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", helloHandler).Methods("POST")
	router.HandleFunc("/api/get-task", getTaskHundler).Methods("GET")
	http.ListenAndServe(":8080", router)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task = req.Message

	fmt.Fprintf(w, "Таска обновлена: %s", task)

}

func getTaskHundler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "НАША ЗАДАЧА ЭТО: %s", task)
}
