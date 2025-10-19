package handlers

import (
	"FoodDeliveryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var menuItems []models.MenuItem

func CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	var menuItem models.MenuItem
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	menuItems = append(menuItems, menuItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menuItem)
}
