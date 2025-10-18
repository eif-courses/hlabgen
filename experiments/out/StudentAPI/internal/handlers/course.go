package handlers

import (
	"StudentAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCourse() {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func GetCourses() {
	// Implementation for fetching courses
}
func UpdateCourse() {
	// Implementation for updating a course
}
func DeleteCourse() {
	// Implementation for deleting a course
}
