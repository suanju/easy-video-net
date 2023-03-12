package contribution

import (
	"Go-Live/logic/contribution/videoSocket"
	"Go-Live/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

// VideoSocket  观看视频建立的socket
func (C Controllers) VideoSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)

	//判断是否创建视频socket房间
	id, _ := strconv.Atoi(ctx.Query("videoID"))
	videoID := uint(id)
	if videoSocket.Severe.VideoRoom[videoID] == nil {
		//无人观看主动创建
		videoSocket.Severe.VideoRoom[videoID] = make(videoSocket.UserMapChannel, 10)
	}
	err := videoSocket.CreateVideoSocket(uid, videoID, ws)
	if err != nil {
		response.ErrorWs(ws, "创建socket失败")
	}
}
