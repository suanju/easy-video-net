package socket

import (
	"Go-Live/consts"
	"Go-Live/global"
	userModel "Go-Live/models/users"
	"Go-Live/proto/pb"
	"Go-Live/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Engine struct {
	//直播间
	LiveRoom map[uint]UserMapChannel

	Register     chan LiveRoomEvent
	Cancellation chan LiveRoomEvent
}

type UserMapChannel map[uint]*UserChannel

//UserChannel 用户信息
type UserChannel struct {
	UserInfo userModel.User
	Socket   *websocket.Conn
	MsgList  chan []byte
}

//LiveRoomEvent 直播间事件 注册 离线
type LiveRoomEvent struct {
	RoomID  uint
	Channel *UserChannel
}

var Severe = &Engine{
	LiveRoom:     make(map[uint]UserMapChannel, 10),
	Register:     make(chan LiveRoomEvent, 10),
	Cancellation: make(chan LiveRoomEvent, 10),
}

// Start 启动服务
func (e *Engine) Start() {
	//为每个用户创建直播间socket
	type userList []userModel.User
	users := new(userList)
	global.Db.Select("id").Find(users)
	for _, v := range *users {
		e.LiveRoom[v.ID] = make(UserMapChannel, 10)
	}
	//监听业务通道信息
	//注册离线事件监听
	for {
		select {
		//注册事件
		case registerMsg := <-e.Register:
			logrus.Infof("注册事件 %v", registerMsg)
			//不存在房间直接推出
			if _, ok := e.LiveRoom[registerMsg.RoomID]; !ok {
				//格式化响应
				message := &pb.Message{
					MsgType: consts.Error,
					Data:    []byte("消息格式错误"),
				}
				res, _ := proto.Marshal(message)
				_ = registerMsg.Channel.Socket.WriteMessage(websocket.BinaryMessage, res)
				return
			}
			//添加成员
			e.LiveRoom[registerMsg.RoomID][registerMsg.Channel.UserInfo.ID] = registerMsg.Channel
			//广播用户上线
			err := serviceOnlineAndOfflineRemind(registerMsg, true)
			if err != nil {
				response.ErrorWs(registerMsg.Channel.Socket, err.Error())
			}
			//给用户发送历史消息
			err = serviceResponseLiveRoomHistoricalBarrage(registerMsg)
			if err != nil {
				response.ErrorWs(registerMsg.Channel.Socket, err.Error())
			}

		case cancellationMsg := <-e.Cancellation:
			logrus.Infof("离线事件 %v", cancellationMsg)
			delete(e.LiveRoom[cancellationMsg.RoomID], cancellationMsg.Channel.UserInfo.ID)
			//广播用户下线
			err := serviceOnlineAndOfflineRemind(cancellationMsg, false)
			if err != nil {
				response.ErrorWs(cancellationMsg.Channel.Socket, err.Error())
			}
		}
	}
}

func CreateSocket(ctx *gin.Context, userId uint, roomID uint, conn *websocket.Conn) (err error) {
	//创建UserChannel
	userChannel := new(UserChannel)
	//绑定ws
	userChannel.Socket = conn
	//绑定用户信息
	user := userModel.User{}
	user.Find(userId)
	userChannel.UserInfo = user
	//防止阻塞
	userChannel.MsgList = make(chan []byte, 10)

	//创建用户
	userLiveEvent := LiveRoomEvent{
		RoomID:  roomID,
		Channel: userChannel,
	}
	Severe.Register <- userLiveEvent

	go userLiveEvent.Read()
	go userLiveEvent.Writer()
	return nil
}
