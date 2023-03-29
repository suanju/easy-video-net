package main

import (
	_ "easy-video-net/global/database/mysql"
	_ "easy-video-net/global/database/redis"
	"easy-video-net/logic/contribution/videoSocket"
	liveSocket "easy-video-net/logic/live/socket"
	"easy-video-net/logic/users/chatByUserSocket"
	"easy-video-net/logic/users/chatSocket"
	"easy-video-net/logic/users/noticeSocket"
	"easy-video-net/router"
	"easy-video-net/utils/testing"
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
	//ces
	router.InitRouter()

}
