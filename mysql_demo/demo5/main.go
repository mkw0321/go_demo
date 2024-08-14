package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User4 struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	//连接mysql
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User4{})
	//创建
	//u1 := User4{Name: "qimi", Age: 18}
	//u2 := User4{Name: "jinzhu", Age: 19}
	//db.Create(&u1)
	//db.Create(&u2)
	//查询
	var user4 User4
	db.First(&user4)                  //根据主键查询第一条数据
	fmt.Printf("user4: %#v\n", user4) //格式化输出
	var users []User4
	//db.Find(&users) //查询所有记录
	//fmt.Printf("users: %#v\n", users)
	//fmt.Printf("\n")
	//db.Where("name = ?", "jinzhu").Find(&users)
	db.Model(&users).Update(User4{Name: "kkvvm", Age: 20})

}
