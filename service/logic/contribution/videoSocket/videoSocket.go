package videoSocket

import (
	"Go-Live/consts"
	receive "Go-Live/interaction/receive/socket"
	userModel "Go-Live/models/users"
	"Go-Live/utils/response"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Engine struct {
	//视频房间
	VideoRoom map[uint]UserMapChannel

	Register     chan VideoRoomEvent
	Cancellation chan VideoRoomEvent
}

type UserMapChannel map[uint]*UserChannel

type ChanInfo struct {
	Type string
	Data interface{}
}

//UserChannel 用户信息
type UserChannel struct {
	UserInfo userModel.User
	Socket   *websocket.Conn
	MsgList  chan ChanInfo
}

//VideoRoomEvent 事件 注册 离线
type VideoRoomEvent struct {
	VideoID uint
	Channel *UserChannel
}

var Severe = &Engine{
	VideoRoom:    make(map[uint]UserMapChannel, 10),
	Register:     make(chan VideoRoomEvent, 10),
	Cancellation: make(chan VideoRoomEvent, 10),
}

// Start 启动服务
func (e *Engine) Start() {
	for {
		select {
		//注册事件
		case registerMsg := <-e.Register:
			//添加成员
			e.VideoRoom[registerMsg.VideoID][registerMsg.Channel.UserInfo.ID] = registerMsg.Channel
			//广播在线观看人数
			num := len(e.VideoRoom[registerMsg.VideoID])
			r := struct {
				People int `json:"people"`
			}{
				People: num,
			}
			res := ChanInfo{
				Type: consts.VideoSocketTypeNumberOfViewers,
				Data: r,
			}
			for _, v := range e.VideoRoom[registerMsg.VideoID] {
				v.MsgList <- res
			}
		case cancellationMsg := <-e.Cancellation:
			//广播在线观看人数
			num := len(e.VideoRoom[cancellationMsg.VideoID]) - 1
			r := struct {
				People int `json:"people"`
			}{
				People: num,
			}
			res := ChanInfo{
				Type: consts.VideoSocketTypeNumberOfViewers,
				Data: r,
			}
			for _, v := range e.VideoRoom[cancellationMsg.VideoID] {
				v.MsgList <- res
			}
			delete(e.VideoRoom[cancellationMsg.VideoID], cancellationMsg.Channel.UserInfo.ID)
		}
	}
}

func CreateVideoSocket(userID uint, videoID uint, conn *websocket.Conn) (err error) {
	//创建UserChannel
	userChannel := new(UserChannel)
	//绑定ws
	userChannel.Socket = conn
	//绑定用户信息
	user := userModel.User{}
	user.Find(userID)
	userChannel.UserInfo = user
	//防止阻塞
	userChannel.MsgList = make(chan ChanInfo, 10)

	//创建用户
	userLiveEvent := VideoRoomEvent{
		VideoID: videoID,
		Channel: userChannel,
	}
	Severe.Register <- userLiveEvent

	go userLiveEvent.Read()
	go userLiveEvent.Writer()
	return nil

}

//Writer 监听写入数据
func (lre VideoRoomEvent) Writer() {
	for {
		select {
		case msg := <-lre.Channel.MsgList:
			response.SuccessWs(lre.Channel.Socket, msg.Type, msg.Data)
		}
	}
}

//Read 读取数据
func (lre VideoRoomEvent) Read() {
	//链接断开进行离线
	defer func() {
		Severe.Cancellation <- lre
		err := lre.Channel.Socket.Close()
		if err != nil {
			return
		}
	}()
	//监听业务通道
	for {
		//检查通达ping通
		lre.Channel.Socket.PongHandler()
		_, text, err := lre.Channel.Socket.ReadMessage()
		if err != nil {
			return
		}
		info := new(receive.Receive)
		if err = json.Unmarshal(text, info); err != nil {
			response.ErrorWs(lre.Channel.Socket, "消息格式错误")
		}
		switch info.Type {

		}
	}
}
