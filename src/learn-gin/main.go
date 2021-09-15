package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func channelTest() {
	wg := sync.WaitGroup{}
	chanel1 := make(chan int, 1)
	chanel2 := make(chan int, 1)
	wg.Add(2)
	go func() {
		chanel1 <- 15
		chanel2 <- 15
		chanel1 <- 15
		chanel2 <- 15
		wg.Done()
	}()
	go func() {
		flag := true
		for flag {
			select {
			case i := <-chanel1:
				fmt.Println(i)
			case i := <-chanel2:
				fmt.Println(i)
			default:
				{
					flag = false
				}
			}

		}
		wg.Done()
	}()
	wg.Wait()
}

//动态参数,
func handle(c *gin.Context) {
	listData := []string{
		"北京",
		"天安门",
		"南京",
		"天津大金矿",
		"天津大麻花",
	}
	data := map[string]interface{}{
		"name": "kiwi",
		"age":  18,
		"list": listData,
	}
	c.JSON(200, data)
}
func applicationRun() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		names := c.QueryArray("name")

		c.JSON(200, gin.H{
			"message": "pong",
			"names":   names,
		})
	}, handle)
	router.GET("/get", func(c *gin.Context) {
		data := struct {
			Data map[string]interface{} `json:"data"`
			Msg  string                 `json:"msg"`
			Code int                    `json:"code"`
		}{
			Data: map[string]interface{}{
				"name":   "kiwi",
				"age":    18,
				"school": "南昌大学",
			},
			Msg:  "登录成功",
			Code: 200,
		}
		c.JSON(200, data)
	})
	//获取路径参数，比如有人输入 “ /user/何东昌/18 ” 这样的路径
	router.GET("/user/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func main() {
	applicationRun()

}
