package live

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/live"
	"Go-Live/logic/live"
	"github.com/gin-gonic/gin"
)

type LivesControllers struct {
	controllers.BaseControllers
}

//GetLiveRoom 获取直播房间
func (lv LivesControllers) GetLiveRoom(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := live.GetLiveRoom(uid)
	lv.Response(ctx, results, err)
}

//GetLiveRoomInfo 获取直播房间信息
func (lv LivesControllers) GetLiveRoomInfo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetLiveRoomInfoReceiveStruct)); err == nil {
		results, err := live.GetLiveRoomInfo(rec, uid)
		lv.Response(ctx, results, err)
	}
}

//GetBeLiveList 获取正在直播的用户
func (lv LivesControllers) GetBeLiveList(ctx *gin.Context) {
	results, err := live.GetBeLiveList()
	lv.Response(ctx, results, err)
}
