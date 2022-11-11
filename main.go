package main

import (
	_ "Go-Live/global/dataBases/Mysql"
	_ "Go-Live/global/dataBases/Redis"
	"Go-Live/router"
	"Go-Live/utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	router.InitRouter()
}
