package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

type (
	// 定义原始的数据库字段
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
	// 处理返回的字段
	transformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

// 连接mysql
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:080121mkwMKW#@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":      http.StatusCreated,
		"msg":         "successfully created",
		"resource_ID": todo.ID,
	})
}

func fetchallTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo
	db.Find(&todos)
	if len(todos) < 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "todos not found", "resource_ID": 0})
		return
	}
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}
		_todos = append(_todos, transformedTodo{item.ID, item.Title, completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "todo not found", "resource_ID": 0})
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	}
	_todo := transformedTodo{todo.ID, todo.Title, completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "todo not found", "resource_ID": 0})
		return
	}
	completed, err := strconv.Atoi(c.PostForm("completed"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid completed value"})
		return
	}
	// 使用Updates方法一次性更新多个字段
	db.Model(&todo).Updates(todoModel{Title: c.PostForm("title"), Completed: completed})
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": " successfully!"})
}

func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "todo not found", "resource_ID": 0})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": " successfully!"})
}

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchallTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id")
		v1.DELETE("/:id")
	}
	r.Run(":9090")

}
