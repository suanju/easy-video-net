package redis

import (
	"Go-Live/global/configRead"
	"fmt"
	"github.com/go-redis/redis"
)

func ReturnsInstance() *redis.Client {
	var err error
	var redisConfig = configRead.Config.RConfig

	// 创建链接
	Db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.IP, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})
	_, err = Db.Ping().Result()
	if err != nil {
		fmt.Printf("redis连接失败:%v \n", err)
	}
	return Db

}
