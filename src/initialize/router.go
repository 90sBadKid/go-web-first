package initialize

import (
	"github.com/gin-gonic/gin"
	"go-web/middleware"
	"go-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()

	group := engine.Group("")
	group.Use(middleware.ExceptionHandler)

	router.RouterGroupApp.InitBookRouter(group)
	router.RouterGroupApp.InitJwtRouter(group)

	return engine
}
