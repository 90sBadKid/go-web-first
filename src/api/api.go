package api

import (
	"go-web/api/book"
	"go-web/api/system"
)

type ApiGroup struct {
	BookApiGroup   book.BookApiGroup
	SystemApiGroup system.SystemApiGroup
}

var ApiGroupApp = new(ApiGroup)
