package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Add logic to save book to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Add logic to retrieve books from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Book{})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Add logic to update book in database
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Add logic to delete book from database
	w.WriteHeader(http.StatusNoContent)
}
