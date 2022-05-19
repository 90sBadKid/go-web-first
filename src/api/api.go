package api

import (
	"go-web/api/book"
)

type ApiGroup struct {
	BookApiGroup book.BookApiGroup
}

var ApiGroupApp = new(ApiGroup)
