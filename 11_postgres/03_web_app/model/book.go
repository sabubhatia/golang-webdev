package model

type Book struct {
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price string `json:"price"`
}

const (
	DBName = "bookstore"
)