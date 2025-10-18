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

// GetMenuItems handles fetching all menu items.
func GetMenuItems() {
	// Implementation for fetching menu items
}

// UpdateMenuItem handles updating an existing menu item.
func UpdateMenuItem() {
	// Implementation for updating a menu item
}

// DeleteMenuItem handles deleting a menu item.
func DeleteMenuItem() {
	// Implementation for deleting a menu item
}
