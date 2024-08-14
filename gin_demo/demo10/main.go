package main

//中间件 处理请求时可以加入自己的钩子函数（Hook）

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// define a hook
func login(c *gin.Context) {
	fmt.Println("login successfully")
	//计时 统计耗时
	start := time.Now()
	c.Next() //调用剩余的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

func main() {
	r := gin.Default()
	r.Use(login) //全局调用hook
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})

	//设置一个film的路由组
	flimGroup := r.Group("/flim", login)
	{
		flimGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/flim/home",
			})
		})
		flimGroup.GET("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/flim/shop",
			})
		})
	}
	r.Run(":8080")
}
