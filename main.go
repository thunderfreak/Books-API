package main

import (
	"fmt"
	"github.com/thunderfreak/rest-api/controller"
	"github.com/thunderfreak/rest-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Router Initialization
	r := mux.NewRouter()
	fmt.Println("Router initialized! Check endpoints")

	// Setting Mock Books Data
	model.SetMockData()

	// Route Handlers / Endpoints
	r.HandleFunc("/", controller.ServeHome).Methods("GET")
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	r.HandleFunc("/books", controller.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
