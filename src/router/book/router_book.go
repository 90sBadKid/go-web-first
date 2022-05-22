package book

import (
	"github.com/gin-gonic/gin"
	"go-web/api"
)

type BookRouter struct{}

func (r *BookRouter) InitBookRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/book")

	bookApi := api.ApiGroupApp.BookApiGroup.BookApi

	routerGroup.POST("/list", bookApi.PageList)
	routerGroup.POST("/add", bookApi.Save)

}
