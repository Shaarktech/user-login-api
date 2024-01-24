package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	initMongoDB() // Initialize MongoDB connection

	// Define API endpoints
	router.HandleFunc("/api/users", createUserHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", updateUserHandler).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
