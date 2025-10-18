package handlers

import (
	"SchoolAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateClass() {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func GetClasses() {
	// Implementation for getting classes
}
func UpdateClass() {
	// Implementation for updating a class
}
func DeleteClass() {
	// Implementation for deleting a class
}
