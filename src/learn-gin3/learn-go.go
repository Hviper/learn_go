package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//使用中间件,Gin框架运行开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
//这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证，权限校验，
//数据分页，记录日志，消耗统计
//中间件：权限管理 AOP切面

//中间件：统计时间的中间件
func HandlerFunc(c *gin.Context) {
	start := time.Now()
	// Process request 调用后续的处理函数，调用后续的链
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)

}

func CentorUser() {
	router := gin.Default()
	router.Use(HandlerFunc)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "kiwi",
			"code": 200,
			"msg":  "登录成功",
		})
	})
	router.Run(":8080")
}

func main() {
	CentorUser()
}
