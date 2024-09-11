package main

//简单测试gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func someGet(c *gin.Context) {
	fmt.Println("someGet")
	c.Writer.WriteString("hello ! This is my web_start")
}

func main() {
	// 创建一个默认的路由引擎
	sever := gin.Default() //逻辑上的服务器  因为可以同时启动多个
	////在启动一个路有引擎
	//go func() {
	//	sever1 := gin.Default()
	//	sever1.Run(":8081")
	//}()
	//静态路由
	sever.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.String(http.StatusOK, "hello world")
	})
	//参数路由
	sever.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello!这是参数路由"+name)
	})
	//通配符路由
	sever.GET("/views/*.html", func(c *gin.Context) {
		c.String(http.StatusOK, "hello! 这是通配符路由"+c.Param(".html"))
	})
	//查询参数 Query
	sever.GET("/order", func(c *gin.Context) {
		oid := c.Query("id")
		c.String(http.StatusOK, "HELLO! 这是查询参数"+oid)
	})

	sever.GET("/upgrade", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "upgrade success",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//端口在配置文件中读取
	sever.Run(":8080") //显示执行端口
}
