package main

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"
)

// Get all books
func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	bookList := []Book{}
	for _, b := range books {
		bookList = append(bookList, b)
	}
	respondWithJSON(w, http.StatusOK, bookList)
}

// Create a new book
func handleCreateBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	b.ID = uuid.New().String()
	books[b.ID] = b
	respondWithJSON(w, http.StatusCreated, b)
}

// Get a single book
func handleGetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, exists := books[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	respondWithJSON(w, http.StatusOK, book)
}

// Update a book
func handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, exists := books[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	var updated Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	updated.ID = id
	books[id] = updated
	respondWithJSON(w, http.StatusOK, updated)
}

// Delete a book
func handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, exists := books[id]; !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	delete(books, id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book deleted"})
}
