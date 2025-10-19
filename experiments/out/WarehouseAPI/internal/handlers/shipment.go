package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

var shipments []models.Shipment

func CreateShipment(w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shipments = append(shipments, shipment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipment)
}

func GetShipments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shipments)
}
func GetShipment(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single shipment
	w.WriteHeader(http.StatusOK)
}
func UpdateShipment(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a shipment
	w.WriteHeader(http.StatusOK)
}
func DeleteShipment(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a shipment
	w.WriteHeader(http.StatusNoContent)
}
