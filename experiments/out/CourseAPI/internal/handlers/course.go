package handlers

import (
	"CourseAPI/internal/models"
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

func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetCourse(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single course
	w.WriteHeader(http.StatusOK)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a course
	w.WriteHeader(http.StatusOK)
}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a course
	w.WriteHeader(http.StatusNoContent)
}
