package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getBooks")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getBook")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range bookList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createBook")
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(reqBody, &book)
	book.ID = strconv.Itoa(len(bookList) + 1)
	bookList = append(bookList, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateBook")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range bookList {
		if item.ID == params["id"] {
			bookList = append(bookList[:index], bookList[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			bookList = append(bookList, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteBook")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range bookList {
		if item.ID == params["id"] {
			bookList = append(bookList[:index], bookList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bookList)
}
