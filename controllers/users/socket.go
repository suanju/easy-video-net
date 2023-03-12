package users

import (
	"Go-Live/logic/users/noticeSocket"
	"Go-Live/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// NoticeSocket  通知socket
func (us UserControllers) NoticeSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)

	err := noticeSocket.CreateNoticeSocket(uid, ws)
	if err != nil {
		response.ErrorWs(ws, "创建socket失败")
	}

}
