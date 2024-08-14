package main

//路由 路由组

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会走这一条处理逻辑 执行匿名函数中func的内容
	//路由 访问index时会执行func匿名函数中的内容
	//GET 获取信息
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // http。Status表示状态码 也可以用200表示但不建议
			"method": "GET",
		})
	})
	//路由组的组 多用于区分不同的业务线和划分API版本 路由组也是可以嵌套的
	//提取出公共前缀创建一个组 shop组
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/home",
			})
		})
		shopGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/oo",
			})
		})
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/xx",
			})
		})
	}
	//POST 创建
	r.POST("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	//update 修改部分数据时
	r.PUT("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//api delete
	r.DELETE("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//Any拥有多种方法
	//r.Any("/home", func(c *gin.Context) {
	//	switch c.Request.Method {
	//	case "GET":
	//		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	//	case "POST":
	//		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	//	case "PUT":
	//		c.JSON(http.StatusOK, gin.H{"method": "PUT"})
	//	case "DELETE":
	//		c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
	//	}
	//})

	r.Run(":9000")
}
