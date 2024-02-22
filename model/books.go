package model

// Book Model
type Book struct {
	ID        string  `json:"id"`
	Isbn      string  `json:"isbn"`
	Title     string  `json:"title"`
	Author    *Author `json:"author"`
	Publisher string  `json:"publisher"`
	Price     int     `json:"price"`
}

// Author Model
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var BookList []Book

func SetMockData() {
	BookList = append(BookList, Book{ID: "1", Isbn: "9780061120084", Title: "To Kill a Mockingbird", Author: &Author{FirstName: "Harper", LastName: "Lee"}, Publisher: "HarperCollins", Price: 450})
	BookList = append(BookList, Book{ID: "2", Isbn: "9780141439587", Title: "Pride and Prejudice", Author: &Author{FirstName: "Jane", LastName: "Austen"}, Publisher: "Penguin Classics", Price: 300})
	BookList = append(BookList, Book{ID: "3", Isbn: "9781400079179", Title: "The Kite Runner", Author: &Author{FirstName: "Khaled", LastName: "Hosseini"}, Publisher: "Riverhead Books", Price: 399})
	BookList = append(BookList, Book{ID: "4", Isbn: "9780679734772", Title: "Norwegian Wood", Author: &Author{FirstName: "Haruki", LastName: "Murakami"}, Publisher: "Vintage", Price: 200})
	BookList = append(BookList, Book{ID: "5", Isbn: "9780739326227", Title: "The Da Vinci Code", Author: &Author{FirstName: "Dan", LastName: "Brown"}, Publisher: "Doubleday", Price: 550})
}
