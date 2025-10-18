package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateShipment handles the creation of a new shipment.
func CreateShipment() {
	var shipment models.Shipment
	if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipment)
}

// GetShipments handles fetching all shipments.
func GetShipments() {
	// Implementation for fetching shipments will go here.
}

// UpdateShipment handles updating an existing shipment.
func UpdateShipment() {
	// Implementation for updating a shipment will go here.
}

// DeleteShipment handles deleting a shipment.
func DeleteShipment() {
	// Implementation for deleting a shipment will go here.
}
