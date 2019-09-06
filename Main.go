package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init books var as a slice Book structure
var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through the books and find with id

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Moc id
	books = append(books, book)                 // Append to the global books
	json.NewEncoder(w).Encode(book)

}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			//Create new
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book) // Append to the global books
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

func main() {

	// Init router
	r := mux.NewRouter()

	// Moc data
	books = append(books, Book{ID: "1", Title: "Book One", Author: &Author{FirstName: "Sanjay", LastName: "shr"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "Sanjay", LastName: "shr"}})
	books = append(books, Book{ID: "3", Title: "Book Three", Author: &Author{FirstName: "Sanjay", LastName: "shr"}})
	books = append(books, Book{ID: "4", Title: "Book Four", Author: &Author{FirstName: "Sanjay", LastName: "shr"}})
	books = append(books, Book{ID: "5", Title: "Book Five", Author: &Author{FirstName: "Sanjay", LastName: "shr"}})

	// Route handlers // endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/book", createBook).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE", "OPTIONS")

	//CORS
	handler := cors.Default().Handler(r)

	// Create server
	log.Fatal(http.ListenAndServe(":8000", handler))
}
