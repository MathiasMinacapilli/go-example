package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

import "go-example/handlers"

func main() {
	// models.CreateModel()
	mux := mux.NewRouter()

	mux.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user", handlers.GetUser).Methods("POST")
	mux.HandleFunc("/api/user/{userId:[0-9]+}", nil).Methods("GET")
	mux.HandleFunc("/api/user/{userId:[0-9]+}", nil).Methods("PUT")
	mux.HandleFunc("/api/user/{userId:[0-9]+}", nil).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

