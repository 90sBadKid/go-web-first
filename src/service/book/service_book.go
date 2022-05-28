package book

import (
	"go-web/global/db"
	"go-web/global/page"
	"go-web/model/book"
	"go-web/model/book/req"
)

type BookService struct{}

func (b *BookService) PageList(pageQuery *req.ModelBookPageQuery) *page.PageModel {
	tx := db.GVA_DB.Where("name LIKE ?", "%"+pageQuery.Name+"%")
	pageModel := page.Paginate(tx, pageQuery.PageQueryModel, book.Book{})
	return pageModel
}

func (b *BookService) Save(book *book.Book) *book.Book {
	db.GVA_DB.Create(book)
	return book
}
