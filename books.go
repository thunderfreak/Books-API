package main

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Book Struct
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Publisher string `json:"publisher"`
	Price int `json:"price"`
}
