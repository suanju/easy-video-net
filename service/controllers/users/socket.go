package users

import (
	"easy-video-net/logic/users/chatByUserSocket"
	"easy-video-net/logic/users/chatSocket"
	"easy-video-net/logic/users/noticeSocket"
	"easy-video-net/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

// NoticeSocket  通知socket
func (us UserControllers) NoticeSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)
	err := noticeSocket.CreateNoticeSocket(uid, ws)
	if err != nil {
		response.ErrorWs(ws, "创建通知socket失败")
	}
}

// ChatSocket  聊天socket
func (us UserControllers) ChatSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)
	err := chatSocket.CreateChatSocket(uid, ws)
	if err != nil {
		response.ErrorWs(ws, "创建聊天socket失败")
	}
}

func (us UserControllers) ChatByUserSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	//判断是否创建视频socket房间
	tidQuery, _ := strconv.Atoi(ctx.Query("tid"))
	tid := uint(tidQuery)
	ws := conn.(*websocket.Conn)
	err := chatByUserSocket.CreateChatByUserSocket(uid, tid, ws)
	if err != nil {
		response.ErrorWs(ws, "创建用户聊天socket失败")
	}
}
