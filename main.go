package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var bookList = []Book{}

func main() {
	// Router Initialization
	r := mux.NewRouter()
	fmt.Println("Router initialized! Check endpoints")
	
	// Mock Data
	bookList = append(bookList, Book{ID: "1", Isbn: "9780061120084", Title: "To Kill a Mockingbird", Author: Author{FirstName: "Harper", LastName: "Lee"}, Publisher: "HarperCollins", Price: 450})
	bookList = append(bookList, Book{ID: "2", Isbn: "9780141439587", Title: "Pride and Prejudice", Author: Author{FirstName: "Jane", LastName: "Austen"}, Publisher: "Penguin Classics", Price: 300})
	bookList = append(bookList, Book{ID: "3", Isbn: "9781400079179", Title: "The Kite Runner", Author: Author{FirstName: "Khaled", LastName: "Hosseini"}, Publisher: "Riverhead Books", Price: 399})
	bookList = append(bookList, Book{ID: "4", Isbn: "9780679734772", Title: "Norwegian Wood", Author: Author{FirstName: "Haruki", LastName: "Murakami"}, Publisher: "Vintage", Price: 200})
	bookList = append(bookList, Book{ID: "5", Isbn: "9780739326227", Title: "The Da Vinci Code", Author: Author{FirstName: "Dan", LastName: "Brown"}, Publisher: "Doubleday", Price: 550})

	// fmt.Println(bookList)

	// Route Handlers / Endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
