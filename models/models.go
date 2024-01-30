package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var MockDataUserLogic = User{
	Email:    "test@mail.com",
	Password: "test1234",
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// initial data books insteads database first
func InitDataBook() []Book {
	books = append(books, Book{
		ID:     1,
		Title:  "Macro",
		Author: "tim jackson",
	})
	books = append(books, Book{
		ID:     2,
		Title:  "Electrical",
		Author: "nana",
	})

	return books
}
