package contribution

import (
	"Go-Live/logic/contribution"
	"Go-Live/models/contribution/video"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

//CreateVideoContribution 发布视频
func (C Controllers) CreateVideoContribution(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	CreateVideoContributionReceive := new(video.CreateVideoContributionReceiveStruct)
	if err := ctx.ShouldBind(CreateVideoContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.CreateVideoContribution(CreateVideoContributionReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
