package main

import (
	"fmt"
	"net/http"

	"github.com/generalchrist/inventory/internal/api"
	"github.com/generalchrist/inventory/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Connect DB
	db.Connect()
	db.Migrate()

	// Router
	r := chi.NewRouter()

	// --- Middleware ---
	r.Use(middleware.RequestID) // adds request ID to context
	r.Use(middleware.Logger)    // logs each request
	r.Use(middleware.Recoverer) // recovers from panics
	r.Use(middleware.RealIP)    // gets client IP from request

	// Routes
	api.RegisterRoutes(r)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
