package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //创建一个默认的路由引擎
	r.LoadHTMLFiles("./gin_demo/demo8/index.html")
	//请求/index路径时，会执行后面的匿名函数

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		//将文件保存到本地（服务端本地）
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将文件保存到本地（服务端本地）
			dst := fmt.Sprint("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":8080")
}
