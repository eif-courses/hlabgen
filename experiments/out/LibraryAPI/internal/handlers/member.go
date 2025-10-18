package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateMember() {
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)
}

func GetMembers() {
	// Implementation for retrieving members
}
func UpdateMember() {
	// Implementation for updating a member
}
func DeleteMember() {
	// Implementation for deleting a member
}
