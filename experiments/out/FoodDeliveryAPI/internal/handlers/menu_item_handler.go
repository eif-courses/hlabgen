package handlers

import (
	"FoodDeliveryAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateMenuItem handles the creation of a new menu item.
func CreateMenuItem() {
	var menuItem models.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menuItem)
}
