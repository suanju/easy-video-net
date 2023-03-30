package contribution

import (
	"easy-video-net/logic/contribution/sokcet"
	"easy-video-net/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

// VideoSocket  观看视频建立的socket
func (c Controllers) VideoSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)
	//判断是否创建视频socket房间
	id, _ := strconv.Atoi(ctx.Query("videoID"))
	videoID := uint(id)
	if sokcet.Severe.VideoRoom[videoID] == nil {
		//无人观看主动创建
		sokcet.Severe.VideoRoom[videoID] = make(sokcet.UserMapChannel, 10)
	}
	err := sokcet.CreateVideoSocket(uid, videoID, ws)
	if err != nil {
		response.ErrorWs(ws, "创建socket失败")
	}
}
