package main

import (
	"log"
	"net/http"
	"os"

	"github.com/d28035203/scaling-waddle/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, relying on process environment")
	}

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "9010"
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	addr := host + ":" + port
	log.Printf("bookstore API listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
