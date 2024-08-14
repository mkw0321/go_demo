package main

import (
	"fmt"
	"github.com/wild-River2016/tdengine_gorm-master"
	"gorm.io/gorm"
	"log"
	"time"
)

type SensorData struct {
	Ts          int     `gorm:"column:ts;`
	Temperature float32 `gorm:"column:temperature;type:float"`
	Humidity    float32 `gorm:"column:humidity;type:float"`
	DeviceID    string  `gorm:"column:device_id;type:NCHAR"`
}

func ConnectTDengine() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@/tcp(%s:%d)/%s?loc=Local", "root", "taosdata", "172.20.6.165", 6030, "datacenter")
	db, err := gorm.Open(tdengine_gorm.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *gorm.DB) error {
	sqlStmt := `
		CREATE TABLE sensor_data (
    ts TIMESTAMP,
    temperature FLOAT,
    humidity FLOAT,
    device_id BINARY(20)
)
	`
	err := db.Exec(sqlStmt).Error
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// 连接到 TDengine
	db, err := ConnectTDengine()
	if err != nil {
		log.Fatal("Error connecting to TDengine: ", err)
	}

	// 创建表
	//err = db.Debug().Table("sensor_data").AutoMigrate(&SensorData{})
	//if err != nil {
	//	log.Fatal("Error creating table: ", err)
	//}
	//err = CreateTable(db)
	//	//if err != nil {
	//	//	return
	//	//}

	// 插入数据
	sensorData := SensorData{
		Ts: int(time.Now().UnixMilli()),
		//Ts:          strconv.FormatInt((time.Now().UnixMilli()), 10),
		Temperature: 23.5,
		Humidity:    60.2,
		DeviceID:    "sensor1",
	}
	err = db.Table("sensor_data").Create(&sensorData).Error
	if err != nil {
		fmt.Println("Error inserting data: ", err)
	}
	var sql = "INSERT INTO sensor_data (ts,temperature,humidity,device_id) VALUES (1720770840002,23.5,60.2,'sensor1')"
	err = db.Exec(sql).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	// 查询数据
	//var data []SensorData
	//err = db.Table("sensor_data").Find(&data).Error
	//if err != nil {
	//	log.Fatal("Error querying data: ", err.Error)
	//}
	//
	//// 打印查询结果
	//for _, row := range data {
	//	fmt.Printf("Timestamp: %s, Temperature: %.2f, Humidity: %.2f, Device ID: %s\n",
	//		row.Ts, row.Temperature, row.Humidity, row.DeviceID)
	//}
}
