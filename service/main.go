package main

import (
	_ "Go-Live/global/database/mysql"
	_ "Go-Live/global/database/redis"
	videoSocket "Go-Live/logic/contribution/videoSocket"
	liveSocket "Go-Live/logic/live/socket"
	"Go-Live/logic/users/chatByUserSocket"
	"Go-Live/logic/users/chatSocket"
	"Go-Live/logic/users/noticeSocket"
	"Go-Live/router"
	"Go-Live/utils/testing"
)

func main() {
	//检查直播服务
	testing.LiveSeverTesting()
	//开启直播和视频socket
	go liveSocket.Severe.Start()
	go videoSocket.Severe.Start()
	go noticeSocket.Severe.Start()
	go chatSocket.Severe.Start()
	go chatByUserSocket.Severe.Start()

	router.InitRouter()

}
