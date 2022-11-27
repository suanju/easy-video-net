package live

import (
	"Go-Live/logic/live/socket"
	"Go-Live/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

func (lv LivesControllers) LiveSocket(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)

	//判断是否创建直播间
	liveRoom, _ := strconv.Atoi(ctx.Query("liveRoom"))
	liveRoomID := uint(liveRoom)
	if socket.Severe.LiveRoom[liveRoomID] == nil {
		response.ErrorWs(ws, "直播未开启")
	}

	err := socket.CreateSocket(ctx, userID, liveRoomID, ws)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
}
