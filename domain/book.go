package domain

import (
	"book-crud/model"
	"book-crud/types"
)

type IBookRepo interface {
	GetBooks(bookID uint) []model.Book
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(bookID uint) error
}

type IBookService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(bookID uint) error
}
