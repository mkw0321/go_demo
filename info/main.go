package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Res struct {
	Client_addr string `gorm:"column:client_addr"`
	State       string `gorm:"column:state"`
	Client_port string `gorm:"column:client_port"`
}

func main() {
	//	db, err := gorm.Open(postgres.New(postgres.Config{
	//		DSN:                  fmt.Sprintf(" host=%s user=%s password=%s port=%s dbname=postgres sslmode=disable TimeZone=Asia/Shanghai", "172.20.4.140", "postgres", "Bjftkj@2023", "5432"),
	//		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//	}), &gorm.Config{})
	//	//db, err := gorm.Open(postgres.Open("postgres://postgres:Bjftkj@2023@172.20.4.140:5432"), &gorm.Config{})
	//	//// 检查是否有错误
	//	//if err != nil {
	//	//	fmt.Println("连接数据库失败：", err)
	//	//	return
	//	//}
	//	if err != nil {
	//		fmt.Printf("连接失败")
	//	}
	//	var result interface{}
	//	db.Raw("SELECT * FROM pg_event_trigger;").Scan(&result)
	//	fmt.Println(result)
	//}
	dsn := "host=172.20.4.140 user=postgres password=Bjftkj@2023 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//var id int
	var res []Res
	//var result [][]map[string]string
	var data [][]map[string]string
	var version string
	//formatInt = strconv.FormatInt(time.Now().UnixMilli(), base)
	//db.Table("table2").Select("id").Where("id=?", 1).Find(&id)
	db.Table("pg_stat_activity").Select(" client_addr,state,client_port").
		Find(&res)
	db.Raw("show server_version").Find(&version)
	for _, connect := range res {
		var dataMap []map[string]string
		dataMap = append(dataMap,
			map[string]string{"name": "client_addr", "val": connect.Client_addr},
			map[string]string{"name": "state", "val": connect.State},
			map[string]string{"name": "client_port", "val": connect.Client_port},
		)
		data = append(data, dataMap)

	}
	fmt.Println(data)
	fmt.Println(version)
	//fmt.Println(id)
	//fmt.Println(result)
}
