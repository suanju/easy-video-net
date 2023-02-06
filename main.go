package main

import (
	_ "Go-Live/global/dataBases/mysql"
	_ "Go-Live/global/dataBases/redis"
	videoSocket "Go-Live/logic/contribution/socket"
	liveSocket "Go-Live/logic/live/socket"
	"Go-Live/router"
	"Go-Live/utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	//开启直播和视频socket
	go liveSocket.Severe.Start()
	go videoSocket.Severe.Start()

	router.InitRouter()

}
