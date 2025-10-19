package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var members []models.Member

func CreateMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	members = append(members, member)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
func GetMember(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single member
	w.WriteHeader(http.StatusOK)
}
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a member
	w.WriteHeader(http.StatusOK)
}
func DeleteMember(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a member
	w.WriteHeader(http.StatusNoContent)
}
