package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateBook handles the creation of a new book.
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
	// Logic to save book to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks handles fetching all books.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch books from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Book{})
}

// UpdateBook handles updating an existing book.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to update book in database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook handles deleting a book.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Logic to delete book from database
	w.WriteHeader(http.StatusNoContent)
}
