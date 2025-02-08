package services

import (
	"net/http"

	"github.com/aimanafzal/CloudX/handlers"
)

func RegisterAuthServiceRoutes() {
	http.HandleFunc("/auth", handlers.AuthHandler)
}
