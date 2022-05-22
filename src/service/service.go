package service

import "go-web/service/book"

type ServiceGroup struct {
	BookServiceGroup book.BookServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
