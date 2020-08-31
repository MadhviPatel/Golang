package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Fetch all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getBooks")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Fetch a book with id
func getBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getBook with id")
	w.Header().Set("Content-Type", "application/json")
	reqParams := mux.Vars(req)
	for _, item := range books {
		if item.ID == reqParams["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//json.NewEncoder(w).Encode(&Book{})
}

//Create a book
func createBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside createBook")
	var book Book
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	//append the newly created book into the books slice
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside update book")
	w.Header().Set("Content-Type", "application/json")
	var book Book
	reqParams := mux.Vars(req)
	for index, item := range books {
		if item.ID == reqParams["id"] {
			books = append(books[:index], books[index+1:]...)
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = reqParams["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside delete book")
	w.Header().Set("Content-Type", "application/json")
	reqParams := mux.Vars(req)
	for index, item := range books {
		if item.ID == reqParams["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init router
	router := mux.NewRouter()

	//Mock the data
	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Book One", Author: &Author{Firstname: "Jonh", Lastname: "Dow"}})

	books = append(books, Book{ID: "2", Isbn: "12534", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	//add endpoints to router
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//set the port
	http.ListenAndServe(":8000", router)

}
