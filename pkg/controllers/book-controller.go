package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/d28035203/scaling-waddle/pkg/models"
	"github.com/d28035203/scaling-waddle/pkg/utils"
	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// GetBook lists all books.
func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	writeJSON(w, http.StatusOK, books)
}

// GetBookById returns one book by id.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

// CreateBook creates a book from the JSON body.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if book.Name == "" || book.Author == "" {
		writeError(w, http.StatusBadRequest, "name and author are required")
		return
	}

	created := book.CreateBook()
	writeJSON(w, http.StatusCreated, created)
}

// UpdateBook updates fields on an existing book.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	payload := &models.Book{}
	if err := utils.ParseBody(r, payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	if payload.Name != "" {
		book.Name = payload.Name
	}
	if payload.Author != "" {
		book.Author = payload.Author
	}
	if payload.Publication != "" {
		book.Publication = payload.Publication
	}

	if err := models.UpdateBook(book); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update book")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

// DeleteBook deletes a book by id.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book := models.DeleteBook(id)
	writeJSON(w, http.StatusOK, book)
}
