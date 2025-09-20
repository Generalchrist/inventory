package api

import (
	"net/http"

	"github.com/generalchrist/inventory/internal/auth"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Inventory API"))
	})

	RegisterAuthRoutes(r)

	// Protected routes
	r.Group(func(protected chi.Router) {
		protected.Use(auth.JWTMiddleware)
		RegisterItemRoutes(protected)
	})
}
