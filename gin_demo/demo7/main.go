package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring  URL后面的内容
func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		username := c.DefaultQuery("username", "xiaowangzi")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":9090")
}
