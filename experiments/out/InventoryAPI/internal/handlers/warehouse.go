package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(warehouse)
}
