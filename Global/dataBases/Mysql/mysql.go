package Mysql

import (
	"Go-Live/Global/configRead"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ReturnsInstance() *gorm.DB {
	var err error

	var mysqlConfig = configRead.Config.SqlConfig
	// 创建链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port, mysqlConfig.Database)
	//dsn := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库链接错误- %v \n", err)
	}
	if Db.Error != nil {
		fmt.Printf("数据库错误- %v \n", Db.Error)
	}
	return Db
}
