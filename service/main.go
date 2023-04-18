package main

import (
	_ "easy-video-net/global/database/mysql" //初始化mysql
	_ "easy-video-net/global/database/redis" //初始化redis
	"easy-video-net/router"
	_ "easy-video-net/utils/socket"  //初始化socket
	_ "easy-video-net/utils/testing" //运行环境检查
)

func main() {
	//注册路由
	router.InitRouter()
}
