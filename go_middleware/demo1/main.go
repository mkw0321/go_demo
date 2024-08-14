package main

//Recovery中间件
import (
	"github.com/gin-gonic/gin"
)

// 这个中间件会在发生 panic 时恢复请求，并返回一个 500 Internal Server Error 响应。
func main() {
	router := gin.Default()

	// 使用 Recovery 中间件
	router.Use(gin.Recovery())

	// 定义处理程序
	router.GET("/", func(c *gin.Context) {
		// 故意触发 panic
		panic("something went wrong")
	})
}
