package initialize

import (
	"github.com/gin-gonic/gin"
	"go-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()

	group := engine.Group("")

	router.RouterGroupApp.InitBookRouter(group)

	return engine
}
