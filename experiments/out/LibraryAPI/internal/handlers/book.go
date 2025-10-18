package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateBook handles the creation of a new book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Save book to database (omitted)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks handles fetching all books.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Fetch books from database (omitted)
	var books []models.Book
	json.NewEncoder(w).Encode(books)
}

// UpdateBook handles updating a book.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Update book logic (omitted)
}

// DeleteBook handles deleting a book.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Delete book logic (omitted)
}
