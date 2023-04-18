package redis

import (
	"easy-video-net/global/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func ReturnsInstance() *redis.Client {
	var err error
	var redisConfig = config.Config.RConfig

	// 创建链接
	Db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.IP, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})
	_, err = Db.Ping().Result()
	if err != nil {
		log.Fatalf("redis连接失败:%v \n", err)
	}
	return Db
}
