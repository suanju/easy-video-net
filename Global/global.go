package Global

import (
	"Go-Live/Global/configRead"
	"Go-Live/Global/dataBases/Mysql"
	RedisDbFun "Go-Live/Global/dataBases/Redis"
	log "Go-Live/Global/logrus"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	Logger = log.ReturnsInstance()
	RedisDb = RedisDbFun.ReturnsInstance()
	Db = Mysql.ReturnsInstance()
	Config = configRead.ReturnsInstance()

}

var (
	Logger  *logrus.Logger
	Config  *configRead.ConfigStruct
	Db      *gorm.DB
	RedisDb *redis.Client
)
