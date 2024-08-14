package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

//var (
//	DB *gorm.DB
//)

type Todo struct {
	ID     int    `json:"id"`
	Tittle string `json:"title"`
	Status bool   `json:"status"`
}

//连接数据库
//func initMySQL() (err error) {
//	dsn := "root:080121mkwMKW#@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local"
//	DB, err = gorm.Open("mysql", dsn)
//	if err != nil {
//		return
//	}
//	return DB.DB().Ping()
//}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	//err := initMySQL()
	if err != nil {
		panic(err)
	}
	//创建默认路由引擎
	r := gin.Default()
	defer db.Close()
	db.AutoMigrate(&Todo{})                                 //绑定
	r.Static("./bubble_project/static", "static")           //模板文件引用的静态文件路径
	r.LoadHTMLGlob("./bubble_project/templates/index.html") //gin框架前端文件路径
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":9090")

	//v1  CRUD
	v1 := r.Group("v1")
	{
		//增加待办事项
		//定义一个/todo的路由
		v1.POST("/todo", func(c *gin.Context) {
			//前端页面填写一个事项并提交
			//1.从请求中取出数据
			//2.传入数据库
			//3.返回响应
			var todo Todo     //定义todo存储传入的数据
			c.BindJSON(&todo) // 解析请求体中的 JSON 数据并绑定到todo即从请求出取出数据
			//将todo数据传入数据库中
			if err = db.Create(&todo).Error; err != nil {
				//若出错则返回JSON响应
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查看表里的所有待办事项
		v1.GET("/todo", func(c *gin.Context) {
			var todolist []Todo
			if err := db.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todolist) //c.JSON(http.Status,todolist)
			}
		})
		//查看某一个待办事项
		//v1.GET("/todo/:id", func(c *gin.Context) {
		//
		//})
		//修改待办事项
		v1.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id is not found"})
				return
			}
			var todo Todo
			if err := db.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除待办事项
		v1.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id is not found"})
				return
			}
			if err = db.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"meg": "successfully deleted"})
			}
		})
	}
}
