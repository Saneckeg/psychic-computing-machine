package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", helloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Здарова, дядя Саша")
}
