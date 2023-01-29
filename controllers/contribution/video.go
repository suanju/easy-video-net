package contribution

import (
	receive "Go-Live/interaction/receive/contribution/video"
	"Go-Live/logic/contribution"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

//CreateVideoContribution 发布视频
func (C Controllers) CreateVideoContribution(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	CreateVideoContributionReceive := new(receive.CreateVideoContributionReceiveStruct)
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

// GetVideoContributionByID  根据id获取视频信息
func (C Controllers) GetVideoContributionByID(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	GetVideoContributionByIDReceive := new(receive.GetVideoContributionByIDReceiveStruct)
	if err := ctx.ShouldBind(GetVideoContributionByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoContributionByID(GetVideoContributionByIDReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
