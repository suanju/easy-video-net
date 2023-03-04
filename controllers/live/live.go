package live

import (
	receive "Go-Live/interaction/receive/live"
	"Go-Live/logic/live"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"

	"github.com/gin-gonic/gin"
)

type LivesControllers struct {
}

//GetLiveRoom 获取直播房间
func (lv LivesControllers) GetLiveRoom(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetLiveRoomReceive := new(receive.GetLiveRoomReceiveStruct)
	if err := ctx.ShouldBind(GetLiveRoomReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := live.GetLiveRoom(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetLiveRoomInfo 获取直播房间信息
func (lv LivesControllers) GetLiveRoomInfo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetLiveRoomInfoReceive := new(receive.GetLiveRoomInfoReceiveStruct)
	if err := ctx.ShouldBind(GetLiveRoomInfoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := live.GetLiveRoomInfo(GetLiveRoomInfoReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetBeLiveList 获取正在直播的用户
func (lv LivesControllers) GetBeLiveList(ctx *gin.Context) {
	results, err := live.GetBeLiveList()
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
