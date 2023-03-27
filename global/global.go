package global

import (
	"Go-Live/global/config"
	"Go-Live/global/database/mysql"
	RedisDbFun "Go-Live/global/database/redis"
	log "Go-Live/global/logrus"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	Logger = log.ReturnsInstance()
	RedisDb = RedisDbFun.ReturnsInstance()
	Db = mysql.ReturnsInstance()
	Config = config.ReturnsInstance()

}

var (
	Logger  *logrus.Logger
	Config  *config.ConfigStruct
	Db      *gorm.DB
	RedisDb *redis.Client
)
