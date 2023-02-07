package main

import (
	"fmt"
	"interview/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := mux.NewRouter()

	r.HandleFunc("/search", handler.SearchHandler).Methods("GET")
	r.HandleFunc("/detail/{id}", handler.DetailHandler).Methods("GET")

	port := 8080
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	log.Printf("Starting server. http port: %d", port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
