package system

import (
	"github.com/gin-gonic/gin"
	"go-web/api"
)

type JwtRouter struct {
}

func (r *JwtRouter) InitJwtRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/jwt")

	jwtApi := api.ApiGroupApp.SystemApiGroup.JwtApi

	routerGroup.POST("/", jwtApi.GenerateToken)

}
