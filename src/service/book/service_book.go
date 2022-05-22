package book

import (
	"go-web/global/db"
	"go-web/global/model"
	"go-web/model/book"
)

type BookService struct{}

func (b *BookService) PageList(page model.PageQueryModel) *model.PageModel {
	pageModel := db.GVA_DB.Page(page, book.Book{})
	return pageModel
}

func (b *BookService) Save(book *book.Book) *book.Book {
	db.GVA_DB.Create(book)
	return book
}
