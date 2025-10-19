package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

var items []models.Item

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func GetItem(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single item
	w.WriteHeader(http.StatusOK)
}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an item
	w.WriteHeader(http.StatusOK)
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an item
	w.WriteHeader(http.StatusNoContent)
}
