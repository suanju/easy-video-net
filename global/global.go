package global

import (
	"Go-Live/global/configRead"
	"Go-Live/global/dataBases/mysql"
	RedisDbFun "Go-Live/global/dataBases/redis"
	log "Go-Live/global/logrus"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	Logger = log.ReturnsInstance()
	RedisDb = RedisDbFun.ReturnsInstance()
	Db = mysql.ReturnsInstance()
	Config = configRead.ReturnsInstance()

}

var (
	Logger  *logrus.Logger
	Config  *configRead.ConfigStruct
	Db      *gorm.DB
	RedisDb *redis.Client
)
