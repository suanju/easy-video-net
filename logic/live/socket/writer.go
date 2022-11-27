package socket

import (
	"github.com/gorilla/websocket"
)

//Writer 监听写入数据
func (lre LiveRoomEvent) Writer() {
	for {
		select {
		case msg := <-lre.Channel.MsgList:
			//fmt.Printf("用户%s 得到消息 %s \n", lre.Channel.UserInfo.Username, msg)
			err := lre.Channel.Socket.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				return
			}
		}
	}
}

////公共消息监听
//func listening() {
//	for {
//		select {
//		//公共群聊
//		case barrageMessage := <-Severe.Barrage:
//			for k, v := range Severe.UserMapChannel {
//				if k != barrageMessage.Sender {
//					v.MsgList <- barrageMessage
//				}
//			}
//		}
//	}
//}
