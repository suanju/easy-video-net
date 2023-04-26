package redis

import (
	"context"
	"easy-video-net/global/config"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sethvargo/go-retry"
	"log"
	"time"
)

var Db *redis.Client

func ReturnsInstance() *redis.Client {
	var err error
	var redisConfig = config.Config.RConfig
	b := retry.NewFibonacci(10 * time.Second)
	ctx := context.Background()
	if err := retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		Db = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", redisConfig.IP, redisConfig.Port),
			Password: redisConfig.Password, // no password set
			DB:       0,                    // use default DB
		})
		_, err = Db.Ping().Result()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		// handle error
		log.Fatalf("重连5次后redis连接失败- %v \n", err)
	}
	return Db
}
