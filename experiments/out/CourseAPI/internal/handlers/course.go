package handlers

import (
	"CourseAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateCourse handles the creation of a new course.
func CreateCourse() {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save course to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

// GetCourses handles fetching all courses.
func GetCourses() {
	// Logic to retrieve courses from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Course{})
}

// UpdateCourse handles updating an existing course.
func UpdateCourse() {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to update course in database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}

// DeleteCourse handles deleting a course.
func DeleteCourse() {
	// Logic to delete course from database
	w.WriteHeader(http.StatusNoContent)
}
