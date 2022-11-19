package main

import (
	_ "Go-Live/global/dataBases/mysql"
	_ "Go-Live/global/dataBases/redis"
	"Go-Live/router"
	"Go-Live/utils/testing"
)

func main() {
	//检查直播服务

	testing.LiveSeverTesting()
	router.InitRouter()
}
