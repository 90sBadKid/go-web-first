package book

import "go-web/service"

type BookApiGroup struct {
	BookApi
}

var (
	bookService = service.ServiceGroupApp.BookServiceGroup.BookService
)
