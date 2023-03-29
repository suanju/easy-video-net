package global

import (
	"easy-video-net/global/config"
	"easy-video-net/global/database/mysql"
	RedisDbFun "easy-video-net/global/database/redis"
	log "easy-video-net/global/logrus"
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
