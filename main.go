package main

import (
	_ "Go-Live/Global/dataBases/Mysql"
	_ "Go-Live/Global/dataBases/Redis"
	"Go-Live/Router"
	"Go-Live/Utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	Router.InitRouter()
}
