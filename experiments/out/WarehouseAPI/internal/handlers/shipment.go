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
	// Implementation for fetching shipments would go here.
}
