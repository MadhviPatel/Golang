package main

import (
	"io/ioutil"
	"net/http"
)

type book struct {
	title  string
	year   string
	id     int
	author *author
}
type author struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	http.HandleFunc("/api/books", getBooks)
	http.HandleFunc("/api/books/book", addBook)

	http.ListenAndServe(":9000", nil)
}

func getBooks(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte("One Book"))
}

func addBook(r http.ResponseWriter, req *http.Request) {
	body := ioutil.ReadAll(req.Body)
	
	r.WriteHeader(200)
}
