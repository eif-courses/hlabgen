package handlers

import (
	"WarehouseAPI/internal/models"
	"encoding/json"
	"net/http"
)

var locations []models.Location

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location models.Location
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	locations = append(locations, location)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(location)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
func GetLocation(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single location
	w.WriteHeader(http.StatusOK)
}
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a location
	w.WriteHeader(http.StatusOK)
}
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a location
	w.WriteHeader(http.StatusNoContent)
}
