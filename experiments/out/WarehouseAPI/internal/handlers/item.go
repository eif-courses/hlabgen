package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateItem handles the creation of a new item.
func CreateItem() {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetItems handles fetching all items.
func GetItems() {
	// Implementation for fetching items would go here.
}
