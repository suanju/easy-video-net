package main

import (
	_ "Go-Live/global/dataBases/mysql"
	_ "Go-Live/global/dataBases/redis"
	"Go-Live/logic/live/socket"
	"Go-Live/router"
	"Go-Live/utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	//开启直播socket
	go socket.Severe.Start()

	router.InitRouter()

}
