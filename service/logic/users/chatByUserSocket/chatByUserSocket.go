package chatByUserSocket

import (
	"Go-Live/consts"
	"Go-Live/global"
	receive "Go-Live/interaction/receive/socket"
	socketResponse "Go-Live/interaction/response/socket"
	"Go-Live/logic/users/chatSocket"
	userModel "Go-Live/models/users"
	"Go-Live/models/users/chat/chatList"
	"Go-Live/models/users/notice"
	"Go-Live/utils/response"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Engine struct {
	//视频房间
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
	Tid      uint
	Socket   *websocket.Conn
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
			//清空未读消息
			cl := new(chatList.ChatsListInfo)
			err := cl.UnreadEmpty(registerMsg.UserInfo.ID, registerMsg.Tid)
			//添加在线记录
			if _, ok := chatSocket.Severe.UserMapChannel[registerMsg.UserInfo.ID]; ok {
				//聊天对象在线
				chatSocket.Severe.UserMapChannel[registerMsg.UserInfo.ID].ChatList[registerMsg.Tid] = registerMsg.Socket
			}
			if err != nil {
				global.Logger.Error("uid %d tid %d 清空未读消息数量失败", registerMsg.UserInfo.ID, registerMsg.Tid)
			}

		case cancellationMsg := <-e.Cancellation:
			//删除成员
			delete(e.UserMapChannel, cancellationMsg.UserInfo.ID)
			//删除在线记录
			if _, ok := chatSocket.Severe.UserMapChannel[cancellationMsg.UserInfo.ID]; ok {
				//聊天对象在线
				delete(chatSocket.Severe.UserMapChannel[cancellationMsg.UserInfo.ID].ChatList, cancellationMsg.Tid)
			}
		}
	}
}

func CreateChatByUserSocket(uid uint, tid uint, conn *websocket.Conn) (err error) {
	//创建UserChannel
	userChannel := new(UserChannel)
	//绑定ws
	userChannel.Socket = conn
	user := &userModel.User{}
	user.Find(uid)
	userChannel.UserInfo = user
	userChannel.Tid = tid
	userChannel.MsgList = make(chan ChanInfo, 10)

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
		case "sendChatMsgText":
			sendChatMsgText(lre, lre.UserInfo.ID, lre.Tid, info)
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
