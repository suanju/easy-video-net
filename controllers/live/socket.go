package live

import (
	"Go-Live/consts"
	"Go-Live/logic/live/socket"
	"Go-Live/proto/pb"
	"Go-Live/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func (lv LivesControllers) LiveSocket(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)

	//判断是否创建直播间
	liveRoom, _ := strconv.Atoi(ctx.Query("liveRoom"))
	liveRoomID := uint(liveRoom)
	if socket.Severe.LiveRoom[liveRoomID] == nil {
		message := &pb.Message{
			MsgType: consts.Error,
			Data:    []byte("直播未开启"),
		}
		res, _ := proto.Marshal(message)
		_ = ws.WriteMessage(websocket.BinaryMessage, res)
		return
	}

	err := socket.CreateSocket(ctx, uid, liveRoomID, ws)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
}
