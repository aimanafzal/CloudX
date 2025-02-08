package handlers

import (
	"context"
	"net/http"

	"google.golang.org/api/idtoken"
)

// AuthHandler authenticates a GCP user
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID token from the request header
	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Verify the ID token
	ctx := context.Background()
	audience := "your-client-id.apps.googleusercontent.com" // Replace with your actual client ID
	payload, err := idtoken.Validate(ctx, idToken, audience)
	if err != nil {
		http.Error(w, "Invalid ID token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Extract user information from the payload
	userID, ok := payload.Claims["sub"].(string)
	if !ok {
		http.Error(w, "Invalid ID token payload", http.StatusUnauthorized)
		return
	}
	email, ok := payload.Claims["email"].(string)
	if !ok {
		http.Error(w, "Invalid ID token payload", http.StatusUnauthorized)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"userID": "` + userID + `", "email": "` + email + `"}`))
}
