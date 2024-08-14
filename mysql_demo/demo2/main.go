package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移（结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	//创建行
	u1 := User{1, "kkvvm", "male", "basketball"}
	db.Create(&u1)
	//查询
	var u User
	db.First(&u)
	fmt.Println(u)
	//更新
	db.Model(&u).Update("hobby", "双色球")
	//删除
	//db.Delete(&u)

}
