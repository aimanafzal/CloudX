package main

import (
	"log"
	"net/http"

	"github.com/aimanafzal/CloudX/cmd/services"
)

func main() {
	services.RegisterAuthServiceRoutes()

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
