package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/global/errors"
	"go-web/global/result"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// MiddleWare 定义中间
func MiddleWare(c *gin.Context) {

	defer failureHandler(c)

	c.Next()
}

func failureHandler(c *gin.Context) {
	t := time.Now()
	fmt.Println("中间件开始执行了")
	if err := recover(); err != nil {
		//打印错误堆栈信息
		log.Printf("panic: %v\n", err.(error).Error())
		debug.PrintStack()
		//封装通用json返回
		if e, ok := err.(*errors.ServerError); ok {
			c.JSON(http.StatusOK, result.FailureStatus(e.Code(), e.Error()))
		} else {
			c.JSON(http.StatusOK, result.Failure(e.Error()))
		}
		//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
		c.Abort()
	}

	t2 := time.Since(t)
	fmt.Printf("中间件执行完毕，耗时：%s \n", t2.String())
}
