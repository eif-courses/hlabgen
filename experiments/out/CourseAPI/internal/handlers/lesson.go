package handlers

import (
	"CourseAPI/internal/models"
	"encoding/json"
	"net/http"
)

var lessons []models.Lesson

func CreateLesson(w http.ResponseWriter, r *http.Request) {
	var lesson models.Lesson
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lessons = append(lessons, lesson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lesson)
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}
func GetLesson(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single lesson
	w.WriteHeader(http.StatusOK)
}
func UpdateLesson(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a lesson
	w.WriteHeader(http.StatusOK)
}
func DeleteLesson(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a lesson
	w.WriteHeader(http.StatusNoContent)
}
