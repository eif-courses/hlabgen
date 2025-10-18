package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateSupplier handles the creation of a new supplier.
func CreateSupplier() {
	var supplier models.Supplier
	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	supplier.ID = primitive.NewObjectID()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(supplier)
}

// GetSuppliers handles fetching all suppliers.
func GetSuppliers() {
	// Implementation for fetching suppliers
}

// UpdateSupplier handles updating an existing supplier.
func UpdateSupplier() {
	// Implementation for updating a supplier
}

// DeleteSupplier handles deleting a supplier.
func DeleteSupplier() {
	// Implementation for deleting a supplier
}
