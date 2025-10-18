package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all books
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a book
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a book
}
