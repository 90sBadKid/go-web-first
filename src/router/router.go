package router

import (
	"go-web/router/book"
	"go-web/router/system"
)

type RouterGroup struct {
	book.BookRouterGroup
	system.SystemRouterGroup
}

var RouterGroupApp = new(RouterGroup)
