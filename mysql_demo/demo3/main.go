package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User2 struct {
	gorm.Model   //内嵌
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_id"` //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func (User2) TableName() string {
	return "demo3"
}

func main() {
	db, err := gorm.Open("mysql", "root:080121mkwMKW#@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移（结构体和数据表进行对应）
	db.AutoMigrate(&User2{})

}
