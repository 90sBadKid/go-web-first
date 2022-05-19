package book

import "go-web/server"

type BookApiGroup struct {
	BookApi
}

var (
	bookService = server.ServiceGroupApp.BookServiceGroup.BookService
)
