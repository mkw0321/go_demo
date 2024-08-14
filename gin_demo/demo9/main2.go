package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //设置默认的路由引擎
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET",
		})
	})
	anything := r.Group("/anything")
	{
		anything.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/anything/home"})
		})

	}
	r.Run(":8080")
}
