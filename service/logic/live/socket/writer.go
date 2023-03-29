package socket

import (
	"github.com/gorilla/websocket"
)

//Writer 监听写入数据
func (lre LiveRoomEvent) Writer() {
	for {
		select {
		case msg := <-lre.Channel.MsgList:
			err := lre.Channel.Socket.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				return
			}
		}
	}
}
