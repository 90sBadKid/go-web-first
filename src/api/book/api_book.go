package book

import (
	"github.com/gin-gonic/gin"
	"go-web/global/model"
	"go-web/global/result"
	"go-web/model/book"
)

type BookApi struct{}

func (api *BookApi) PageList(c *gin.Context) {
	var pageQuery model.PageQueryModel
	_ = c.ShouldBind(&pageQuery)

	pageModel := bookService.PageList(pageQuery)
	c.JSON(200, result.SuccessfulData(pageModel))
}

func (api *BookApi) Save(c *gin.Context) {
	var reqBook book.Book
	_ = c.ShouldBind(&reqBook)

	respBook := bookService.Save(&reqBook)
	c.JSON(200, result.SuccessfulData(respBook))
}
