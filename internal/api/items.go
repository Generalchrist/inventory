package api

import (
	"encoding/json"
	"net/http"

	"github.com/generalchrist/inventory/internal/db"
	"github.com/generalchrist/inventory/internal/models"
	"github.com/go-chi/chi/v5"
)

// Register item-related routes
func RegisterItemRoutes(r chi.Router) {
	r.Get("/items", getItems)
	r.Post("/items", createItem)
}

// GET /items
func getItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	err := db.DB.Select(&items, "SELECT * FROM items")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// POST /items
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result, err := db.DB.Exec("INSERT INTO items (name, quantity, location_id) VALUES (?, ?, ?)",
		newItem.Name, newItem.Quantity, newItem.LocationID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, _ := result.LastInsertId()
	newItem.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}
