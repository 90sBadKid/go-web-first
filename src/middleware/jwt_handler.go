package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/global/jwt"
	"go-web/global/result"
	"log"
	"net/http"
)

// JwtHandler 定义中间
func JwtHandler(c *gin.Context) {
	jwtStr := c.GetHeader("auth")
	_, flag, err := jwt.ParseToken(jwtStr)
	if err != nil {
		log.Printf("parse jwt token error! %v \n", err.Error())
	}
	if !flag {
		c.JSON(http.StatusOK, result.FailureStatus(result.StatusTokenOverdueError, result.StatusText(result.StatusTokenOverdueError)))
		c.Abort()
	}
	c.Next()
}
