package mysql

import (
	"easy-video-net/global/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ReturnsInstance() *gorm.DB {
	var err error

	var mysqlConfig = config.Config.SqlConfig
	// 创建链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port, mysqlConfig.Database)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info), 打印所有sql
	})
	if err != nil {
		fmt.Printf("数据库链接错误- %v \n", err)
	}
	if Db.Error != nil {
		fmt.Printf("数据库错误- %v \n", Db.Error)
	}
	return Db
}
