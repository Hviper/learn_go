package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormObj struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//http模块的使用
func UserHttpMethods() {
	router := gin.Default()
	router.GET("/http", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}

func test() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var form FormObj
		//绑定（ShouldBind）需要修改form的值，因此需要将form的指针转递进去
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "成功"})
				fmt.Printf("%#v", form)
				fmt.Println(form.Password)
				fmt.Println(form.User)
			}
		} else {
			c.JSON(404, gin.H{"status": "失败"})
		}

	})

	r.POST("/ping", func(c *gin.Context) {
		var form FormObj
		//绑定（ShouldBind）需要修改form的值，因此需要将form的指针转递进去
		if c.ShouldBind(&form) == nil {
			fmt.Printf("%#v", form)
			c.JSON(200, gin.H{"status": "成功", "code": 200})
		} else {
			c.JSON(404, gin.H{"status": "失败", "code": 404})
		}

	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

//处理函数
func handler(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"name": "kiwi",
		"age":  18,
	})
}
func RouterGroup() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/api")
	{
		v1.GET("/login", handler)
		v1.GET("/submit", handler)
		v1.GET("/read", handler)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", handler)
		v2.POST("/submit", handler)
		v2.POST("/read", handler)
	}

	router.Run(":8080")
}
func main() {
	RouterGroup()
}
