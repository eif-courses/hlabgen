package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var suppliers []models.Supplier

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	suppliers = append(suppliers, supplier)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suppliers)
}
