package handlers

import (
	"HealthAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateRecord handles the creation of a new medical record.
func CreateRecord() {
	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save record to the database goes here.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}

// GetRecords handles fetching all records with pagination.
func GetRecords() {
	// Logic to fetch records from the database goes here.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Record{})
}
