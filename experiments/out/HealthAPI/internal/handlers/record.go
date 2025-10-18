package handlers

import (
	"HealthAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateRecord() {
	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}

func GetRecords() {
	// Implementation for retrieving records
}
func UpdateRecord() {
	// Implementation for updating a record
}
func DeleteRecord() {
	// Implementation for deleting a record
}
