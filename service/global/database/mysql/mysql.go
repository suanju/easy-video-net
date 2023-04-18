package mysql

import (
	"easy-video-net/global/config"
	globalLog "easy-video-net/global/logrus"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type MyWriter struct {
	log *logrus.Logger
}

// Printf 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	m.log.Errorf(fmt.Sprintf(format, v...))
}

func NewMyWriter() *MyWriter {
	log := globalLog.ReturnsInstance()
	return &MyWriter{log: log}
}

func ReturnsInstance() *gorm.DB {
	var err error
	var mysqlConfig = config.Config.SqlConfig
	//sql日志记录
	myLogger := logger.New(
		//设置Logger
		NewMyWriter(),
		logger.Config{
			LogLevel: logger.Error,
		},
	)
	// 创建链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port, mysqlConfig.Database)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: myLogger,
	})
	if err != nil {
		log.Fatalf("数据库链接错误- %v \n", err)
	}
	if Db.Error != nil {
		log.Fatalf("数据库错误- %v \n", Db.Error)
	}
	return Db
}
