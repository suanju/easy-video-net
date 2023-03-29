package chatSocket

import (
	"easy-video-net/consts"
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/socket"
	socketResponse "easy-video-net/interaction/response/socket"
	userLogic "easy-video-net/logic/users"
	userModel "easy-video-net/models/users"
	"easy-video-net/models/users/chat/chatList"
	"easy-video-net/models/users/notice"
	"easy-video-net/utils/response"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Engine struct {
	UserMapChannel map[uint]*UserChannel

	Register     chan *UserChannel
	Cancellation chan *UserChannel
}

type ChanInfo struct {
	Type string
	Data interface{}
}

//UserChannel 用户信息
type UserChannel struct {
	UserInfo *userModel.User
	Socket   *websocket.Conn
	ChatList map[uint]*websocket.Conn
	MsgList  chan ChanInfo
}

var Severe = &Engine{
	UserMapChannel: make(map[uint]*UserChannel, 10),
	Register:       make(chan *UserChannel, 10),
	Cancellation:   make(chan *UserChannel, 10),
}

// Start 启动服务
func (e *Engine) Start() {
	for {
		select {
		//注册事件
		case registerMsg := <-e.Register:
			//添加成员
			e.UserMapChannel[registerMsg.UserInfo.ID] = registerMsg
			//进行未读消息通知
			cl := new(chatList.ChatsListInfo)
			unreadNum := cl.GetUnreadNumber(registerMsg.UserInfo.ID)
			if *unreadNum > 0 {
				//存在未读消息 直接推送聊天列表和记录
				list, err := userLogic.GetChatList(registerMsg.UserInfo.ID)
				if err != nil {
					fmt.Println("查询错误")
					return
				}
				response.SuccessWs(registerMsg.Socket, consts.ChatOnlineUnreadNotice, list)
			}
		case cancellationMsg := <-e.Cancellation:
			//删除成员
			delete(e.UserMapChannel, cancellationMsg.UserInfo.ID)
		}
	}
}

func CreateChatSocket(uid uint, conn *websocket.Conn) (err error) {
	//创建UserChannel
	userChannel := new(UserChannel)
	//绑定ws
	userChannel.Socket = conn
	user := &userModel.User{}
	user.Find(uid)
	userChannel.UserInfo = user
	userChannel.MsgList = make(chan ChanInfo, 10)
	userChannel.ChatList = make(map[uint]*websocket.Conn, 0)

	Severe.Register <- userChannel

	go userChannel.Read()
	go userChannel.Writer()
	return nil

}

//Writer 监听写入数据
func (lre *UserChannel) Writer() {
	for {
		select {
		case msg := <-lre.MsgList:
			response.SuccessWs(lre.Socket, msg.Type, msg.Data)
		}
	}
}

//Read 读取数据
func (lre *UserChannel) Read() {
	//链接断开进行离线
	defer func() {
		Severe.Cancellation <- lre
		err := lre.Socket.Close()
		if err != nil {
			return
		}
	}()
	//监听业务通道
	for {
		//检查通达ping通
		lre.Socket.PongHandler()
		_, text, err := lre.Socket.ReadMessage()
		if err != nil {
			return
		}
		info := new(receive.Receive)
		if err = json.Unmarshal(text, info); err != nil {
			response.ErrorWs(lre.Socket, "消息格式错误")
		}
		switch info.Type {

		}
	}
}

func (lre *UserChannel) NoticeMessage(tp string) {
	//获取未读消息
	nl := new(notice.Notice)
	num := nl.GetUnreadNum(lre.UserInfo.ID)
	if num == nil {
		global.Logger.Error("通知id为%d用户未读消息失败", lre.UserInfo.ID)
	}
	lre.MsgList <- ChanInfo{
		Type: consts.NoticeSocketTypeMessage,
		Data: socketResponse.NoticeMessageStruct{
			NoticeType: tp,
			Unread:     num,
		},
	}
}
