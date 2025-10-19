package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

var courses []models.Course

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	courses = append(courses, course)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}
