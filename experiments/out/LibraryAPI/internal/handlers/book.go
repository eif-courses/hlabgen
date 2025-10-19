package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var books []models.Book

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
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single book
	w.WriteHeader(http.StatusOK)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a book
	w.WriteHeader(http.StatusOK)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a book
	w.WriteHeader(http.StatusNoContent)
}
