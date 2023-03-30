package socket

import (
	"easy-video-net/logic/contribution/sokcet"
	liveSocket "easy-video-net/logic/live/socket"
	"easy-video-net/logic/users/chat"
	"easy-video-net/logic/users/chatUser"
	"easy-video-net/logic/users/notice"
)

func init() {
	//初始化所有socket
	go liveSocket.Severe.Start()
	go sokcet.Severe.Start()
	go notice.Severe.Start()
	go chat.Severe.Start()
	go chatUser.Severe.Start()
}
