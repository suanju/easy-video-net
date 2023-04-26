package mysql

import (
	"context"
	"easy-video-net/global/config"
	globalLog "easy-video-net/global/logrus"
	"fmt"
	"github.com/sethvargo/go-retry"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var Db *gorm.DB

type MyWriter struct {
	log *logrus.Logger
}

// Printf 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	m.log.Errorf(fmt.Sprintf(format, v...))
}

func NewMyWriter() *MyWriter {
	instance := globalLog.ReturnsInstance()
	return &MyWriter{log: instance}
}

func ReturnsInstance() *gorm.DB {
	var mysqlConfig = config.Config.SqlConfig
	//sql日志记录
	myLogger := logger.New(
		//设置Logger
		NewMyWriter(),
		logger.Config{
			LogLevel: logger.Error,
		},
	)
	b := retry.NewFibonacci(10 * time.Second)
	//var Db *gorm.DB
	ctx := context.Background()
	if err := retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		// 创建链接
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port, mysqlConfig.Database)
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: myLogger,
		})
		if err != nil {
			return err
		}
		if Db.Error != nil {
			return err
		}
		return nil
	}); err != nil {
		// handle error
		log.Fatalf("重连3次后依旧数据库链接错误- %v \n", err)
	}
	return Db
}
