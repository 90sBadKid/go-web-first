package book

import (
	"go-web/global/db"
	"go-web/global/model"
	"go-web/model/book"
	"go-web/model/book/req"
)

type BookService struct{}

func (b *BookService) PageList(page req.ModelBookPageQuery) *model.PageModel {
	tx := db.GVA_DB.Where("name LIKE ?", "%"+page.Name+"%")
	pageModel := db.GVA_DB.Page(tx, page.PageQueryModel, book.Book{})
	return pageModel
}

func (b *BookService) Save(book *book.Book) *book.Book {
	db.GVA_DB.Create(book)
	return book
}
