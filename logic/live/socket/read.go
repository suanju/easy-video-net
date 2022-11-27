package socket

import (
	"Go-Live/consts"
	"Go-Live/proto/pb"
	"Go-Live/utils/response"
	"fmt"
	"google.golang.org/protobuf/proto"
)

//Read 读取数据
func (lre LiveRoomEvent) Read() {
	//链接断开进行离线
	defer func() {
		Severe.Cancellation <- lre
		err := lre.Channel.Socket.Close()
		if err != nil {
			return
		}
	}()
	//监听业务通道

	//消息读取
	for {
		//检查通达ping通
		lre.Channel.Socket.PongHandler()
		_, text, err := lre.Channel.Socket.ReadMessage()
		if err != nil {
			return
		}
		data := &pb.Message{}
		if err := proto.Unmarshal(text, data); err != nil {
			response.ErrorWs(lre.Channel.Socket, "消息格式错误")
			fmt.Println("消息格式错误")
		}
		//得到标准格式进行转发
		err = getTypeCorrespondingFunc(lre, data)
		if err != nil {
			response.ErrorWs(lre.Channel.Socket, err.Error())
		}
	}
}
func getTypeCorrespondingFunc(lre LiveRoomEvent, data *pb.Message) error {
	switch data.MsgType {
	case consts.WebClientBarrageReq:
		return serviceSendBarrage(lre, data.Data)
	}
	return fmt.Errorf("消息类型未定义")
}
