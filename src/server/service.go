package server

import "go-web/server/book"

type ServiceGroup struct {
	BookServiceGroup   book.BookServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
