package main

import (
	_ "easy-video-net/global/database/mysql"
	_ "easy-video-net/global/database/redis"
	"easy-video-net/router"
	_ "easy-video-net/utils/socket"
	"easy-video-net/utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	//ces
	router.InitRouter()

}
