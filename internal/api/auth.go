package api

import (
	"encoding/json"
	"net/http"

	"github.com/generalchrist/inventory/internal/auth"
	"github.com/go-chi/chi/v5"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAuthRoutes(r chi.Router) {
	r.Post("/login", loginHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//TODO: Hardcoded user for now
	if creds.Username != "admin" || creds.Password != "password" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := auth.GenerateToken(creds.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
