package socket

import (
	"easy-video-net/consts"
	"easy-video-net/proto/pb"
	"easy-video-net/utils/response"
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
			response.ErrorWsProto(lre.Channel.Socket, "消息格式错误")
		}
		//得到标准格式进行转发
		err = getTypeCorrespondingFunc(lre, data)
		if err != nil {
			response.ErrorWsProto(lre.Channel.Socket, err.Error())
		}
	}
}
func getTypeCorrespondingFunc(lre LiveRoomEvent, data *pb.Message) error {
	switch data.MsgType {
	case consts.WebClientBarrageReq:
		return serviceSendBarrage(lre, data.Data)
	}
	response.ErrorWsProto(lre.Channel.Socket, "未定义的消息格式")
	return nil
}
