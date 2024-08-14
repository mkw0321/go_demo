package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User3 struct {
	ID   int64
	Name string `gorm:"default:'kkvvm'"`
	Age  int64
}

type User4 struct {
	ID   int64
	Name string `gorm:"default:'kkvvm'"`
	Age  int64
}

func main() {
	//连接mysql
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User3{})
	//创建
	u1 := User3{Name: "qimi", Age: 18}

	db.Create(&u1)

}
