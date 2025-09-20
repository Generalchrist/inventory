package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/generalchrist/inventory/internal/db"
	"github.com/generalchrist/inventory/internal/models"
)

// Register a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", 500)
		return
	}

	// insert user
	_, err = db.DB.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, input.Username, string(hash))
	if err != nil {
		http.Error(w, "User already exists or DB error", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created"})
}

// Login a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user models.User
	err := db.DB.Get(&user, `SELECT * FROM users WHERE username = ?`, input.Username)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", 401)
		return
	} else if err != nil {
		http.Error(w, "DB error", 500)
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	// for now: just return success
	json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
}
