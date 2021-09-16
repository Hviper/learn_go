package main

import (
	"github.com/gin-gonic/gin"
)

type Users struct {
	Username string `gorm:"column:mg_name"`
	Password string `gorm:"column:mg_pwd"`
}

func RouterGroup() {
	router := gin.Default()
	list := Query()
	gp := router.Group("/api/private/v1/")
	{
		gp.GET("/login", func(c *gin.Context) {
			var users []Users
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "成功",
				"data": Login(&users),
			})
		})

		gp.GET("/api", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "成功",
				"data": list,
			})
		})
	}

	router.Run()
}

func main() {

	RouterGroup()

}
