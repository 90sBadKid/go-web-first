package system

import (
	"github.com/gin-gonic/gin"
	"go-web/global/jwt"
	"go-web/global/result"
	"net/http"
)

type JwtApi struct{}

func (j *JwtApi) GenerateToken(c *gin.Context) {
	token, _ := jwt.GenerateToken()
	c.JSON(http.StatusOK, result.SuccessfulData(token))
}
