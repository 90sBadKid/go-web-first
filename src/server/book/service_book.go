package book

import (
	"fmt"
	"go-web/global"
	"go-web/model/book"
)

type BookService struct{}

func (b *BookService) GetBookList() []book.Book {
	var bookList = make([]book.Book, 2)
	bookList[0] = book.Book{Name: "西游记", Author: "吴承恩", ChiefEditor: "wry"}
	bookList[1] = book.Book{Name: "三国演义", Author: "罗贯中", ChiefEditor: "wry"}
	return bookList
}

func (b *BookService) AddBook(book *book.Book) (*book.Book, error) {
	tx := global.GVA_DB.Create(book)

	if tx.Error == nil {
		fmt.Println("创建成功！")
		return book, nil
	}
	fmt.Println("创建失败")
	return nil, tx.Error
}
