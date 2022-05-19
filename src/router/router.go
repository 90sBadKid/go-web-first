package router

import "go-web/router/book"

type RouterGroup struct {
	book.BookRouterGroup
}

var RouterGroupApp = new(RouterGroup)
