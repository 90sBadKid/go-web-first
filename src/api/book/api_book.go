package book

import (
	"github.com/gin-gonic/gin"
	"go-web/model/book"
)

type BookApi struct{}

func (api *BookApi) GetBookList(c *gin.Context) {
	list := bookService.GetBookList()
	c.JSON(200, gin.H{
		"message": "成功",
		"data":    list,
	})
}

func (api *BookApi) AddBook(c *gin.Context) {
	var reqBook book.Book
	_ = c.ShouldBind(&reqBook)

	respBook, _ := bookService.AddBook(&reqBook)
	c.JSON(200, gin.H{
		"message": "成功",
		"data":    respBook,
	})
}
