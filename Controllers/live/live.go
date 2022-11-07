package live

import (
	"Go-Live/Logic/live"
	liveModel "Go-Live/Models/live"
	"Go-Live/Utils/response"
	"Go-Live/Utils/validator"
	"github.com/gin-gonic/gin"
)

type LivesControllers struct {
}

//GetLiveRoom 获取直播房间
func (lv LivesControllers) GetLiveRoom(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	getLiveRoomReceive := new(liveModel.GetLiveRoomReceiveStruct)
	if err := ctx.ShouldBind(getLiveRoomReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := live.GetLiveRoom(getLiveRoomReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
