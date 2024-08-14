package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Stu struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func main() {
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var list []Stu
	err = db.Table("mytable").Find(&list).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
