package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateLocation handles the creation of a new location.
func CreateLocation() {
	var location models.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(location)
}

// GetLocations handles fetching all locations.
func GetLocations() {
	// Implementation for fetching locations will go here.
}

// UpdateLocation handles updating an existing location.
func UpdateLocation() {
	// Implementation for updating a location will go here.
}

// DeleteLocation handles deleting a location.
func DeleteLocation() {
	// Implementation for deleting a location will go here.
}
