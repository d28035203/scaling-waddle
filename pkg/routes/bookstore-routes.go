package routes

import (
	"github.com/d28035203/scaling-waddle/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookStoreRoutes registers the routes for the book store CRUD API
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               // Get all books
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // Get book by ID
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           // Create a new book
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // Update a book by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // Delete a book by ID
}
