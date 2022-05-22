package initialize

import (
	"github.com/gin-gonic/gin"
	"go-web/middleware"
	"go-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()

	group := engine.Group("")

	group.Use(middleware.MiddleWare)

	router.RouterGroupApp.InitBookRouter(group)

	return engine
}
