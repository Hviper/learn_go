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
//中间件之间的通信：用全局变量c，c.set(),c.get()方式进行存放和读取值

//中间件：统计时间的中间件
func HandlerFunc(c *gin.Context) {
	start := time.Now()
	// Process request 调用后续的处理函数，调用后续的链
	c.Next()
	// c.Abort()   //终止后续链操作，直接执行我这个链当前的业务操作
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

//传递参数的方式
//闭包传参的形式
func returnFunc(param interface{}) func(*gin.Context) {
	//查询数据库操作
	//或其他准备工作
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		fmt.Printf("cost:%v\n", cost)
		fmt.Printf("param:%v\n", param)
	}
}

//闭包传参的形式
func returnFunc2(param interface{}) gin.HandlerFunc {
	//查询数据库操作
	//或其他准备工作
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		fmt.Printf("cost:%v\n", cost)
		fmt.Printf("param:%v\n", param)
	}
}

func CentorUser() {
	router := gin.Default()
	router.Use(HandlerFunc)
	router.GET("/ping", returnFunc(map[string]string{"name": "kiwi"}), func(c *gin.Context) {
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
