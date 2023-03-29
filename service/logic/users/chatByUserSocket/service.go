package chatByUserSocket

import (
	"easy-video-net/consts"
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/socket"
	"easy-video-net/interaction/response/socket"
	"easy-video-net/logic/users/chatSocket"
	"easy-video-net/models/users/chat/chatList"
	"easy-video-net/models/users/chat/chatMsg"
	"easy-video-net/utils/conversion"
	"easy-video-net/utils/response"
)

func sendChatMsgText(ler *UserChannel, uid uint, tid uint, info *receive.Receive) {

	//添加消息记录
	cm := chatMsg.Msg{
		Uid:     uid,
		Tid:     tid,
		Type:    info.Type,
		Message: info.Data,
	}
	err := cm.AddMessage()
	if err != nil {
		response.ErrorWs(ler.Socket, "发送失败")
		return
	}
	//消息查询
	msgInfo := new(chatMsg.Msg)
	err = msgInfo.FindByID(cm.ID)
	if err != nil {
		response.ErrorWs(ler.Socket, "发送消息失败")
		return
	}
	photo, _ := conversion.FormattingJsonSrc(msgInfo.UInfo.Photo)

	//给自己发消息不推送
	if uid == tid {
		return
	}

	if _, ok := chatSocket.Severe.UserMapChannel[tid]; ok {
		//在线情况
		if _, ok := chatSocket.Severe.UserMapChannel[tid].ChatList[uid]; ok {
			//在与自己聊天窗口 (直接进行推送)
			response.SuccessWs(chatSocket.Severe.UserMapChannel[tid].ChatList[uid], consts.ChatSendTextMsg, socket.ChatSendTextMsgStruct{
				ID:        msgInfo.ID,
				Uid:       msgInfo.Uid,
				Username:  msgInfo.UInfo.Username,
				Photo:     photo,
				Tid:       msgInfo.Tid,
				Message:   msgInfo.Message,
				Type:      msgInfo.Type,
				CreatedAt: msgInfo.CreatedAt,
			})
			return
		} else {
			//添加未读记录
			cl := new(chatList.ChatsListInfo)
			err := cl.UnreadAutocorrection(tid, uid)
			if err != nil {
				global.Logger.Error("uid %d tid %d 消息记录自增未读消息数量失败", tid, uid)
			}
			ci := new(chatList.ChatsListInfo)
			_ = ci.FindByID(uid, tid)
			//推送主socket
			response.SuccessWs(chatSocket.Severe.UserMapChannel[tid].Socket, consts.ChatUnreadNotice, socket.ChatUnreadNoticeStruct{
				Uid:         uid,
				Tid:         tid,
				LastMessage: ci.LastMessage,
				LastMessageInfo: socket.ChatSendTextMsgStruct{
					ID:        msgInfo.ID,
					Uid:       msgInfo.Uid,
					Username:  msgInfo.UInfo.Username,
					Photo:     photo,
					Tid:       msgInfo.Tid,
					Message:   msgInfo.Message,
					Type:      msgInfo.Type,
					CreatedAt: msgInfo.CreatedAt,
				},
				Unread: cl.Unread,
			})
		}
	} else {
		//不在线
		cl := new(chatList.ChatsListInfo)
		err := cl.UnreadAutocorrection(tid, uid)
		if err != nil {
			global.Logger.Error("uid %d tid %d 消息记录自增未读消息数量失败", tid, uid)
		}
	}
}
