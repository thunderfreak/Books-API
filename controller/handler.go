package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thunderfreak/rest-api/model"
	"net/http"
	"strconv"
)

// Controllers
// Home handler
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Books Rest API"))
}

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getBooks - gets all books")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.BookList)
}

// Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getBook - gets one book")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range model.BookList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID")
}

// Create New Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createBook - creates one book")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter data to create book")
	}

	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(reqBody, &book)
	book.ID = strconv.Itoa(len(model.BookList) + 1)
	model.BookList = append(model.BookList, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateBook - updates one book")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["id"] == "" {
		json.NewEncoder(w).Encode("Enter ID of book you want to update in path")
	}

	for index, item := range model.BookList {
		if item.ID == params["id"] {
			model.BookList = append(model.BookList[:index], model.BookList[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			model.BookList = append(model.BookList, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteBook - deletes one book")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range model.BookList {
		if item.ID == params["id"] {
			model.BookList = append(model.BookList[:index], model.BookList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(model.BookList)
}
