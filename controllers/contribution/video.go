package contribution

import (
	receive "Go-Live/interaction/receive/contribution/video"
	"Go-Live/logic/contribution"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controllers struct {
}

//CreateVideoContribution 发布视频和编辑视频
func (C Controllers) CreateVideoContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	CreateVideoContributionReceive := new(receive.CreateVideoContributionReceiveStruct)
	if err := ctx.ShouldBind(CreateVideoContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.CreateVideoContribution(CreateVideoContributionReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//UpdateVideoContribution 编辑视频
func (C Controllers) UpdateVideoContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	UpdateVideoContributionReceive := new(receive.UpdateVideoContributionReceiveStruct)
	if err := ctx.ShouldBind(UpdateVideoContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.UpdateVideoContribution(UpdateVideoContributionReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

// GetVideoContributionByID  根据id获取视频信息
func (C Controllers) GetVideoContributionByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetVideoContributionByIDReceive := new(receive.GetVideoContributionByIDReceiveStruct)
	if err := ctx.ShouldBind(GetVideoContributionByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoContributionByID(GetVideoContributionByIDReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

// SendVideoBarrage  发送视频弹幕
func (C Controllers) SendVideoBarrage(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	SendVideoBarrageReceive := new(receive.SendVideoBarrageReceiveStruct)
	if err := ctx.ShouldBindBodyWith(SendVideoBarrageReceive, binding.JSON); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	_, err := contribution.SendVideoBarrage(SendVideoBarrageReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Default(ctx)
}

// GetVideoBarrage  获取视频弹幕 (播放器）
func (C Controllers) GetVideoBarrage(ctx *gin.Context) {
	GetVideoBarrageReceive := new(receive.GetVideoBarrageReceiveStruct)
	GetVideoBarrageReceive.ID = ctx.Query("id")
	results, err := contribution.GetVideoBarrage(GetVideoBarrageReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.BarrageSuccess(ctx, results)
}

// GetVideoBarrageList  获取视频弹幕展示
func (C Controllers) GetVideoBarrageList(ctx *gin.Context) {
	GetVideoBarrageReceive := new(receive.GetVideoBarrageListReceiveStruct)
	GetVideoBarrageReceive.ID = ctx.Query("id")
	results, err := contribution.GetVideoBarrageList(GetVideoBarrageReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.BarrageSuccess(ctx, results)
}

//VideoPostComment 视频评论
func (C Controllers) VideoPostComment(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	VideoPostCommentReceive := new(receive.VideosPostCommentReceiveStruct)
	if err := ctx.ShouldBind(VideoPostCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.VideoPostComment(VideoPostCommentReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetVideoComment 获取视频评论
func (C Controllers) GetVideoComment(ctx *gin.Context) {
	GetVideoCommentReceive := new(receive.GetVideoCommentReceiveStruct)
	if err := ctx.ShouldBind(GetVideoCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoComment(GetVideoCommentReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetVideoManagementList 创作中心获取视频稿件列表
func (C Controllers) GetVideoManagementList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetVideoManagementListReceive := new(receive.GetVideoManagementListReceiveStruct)
	if err := ctx.ShouldBind(GetVideoManagementListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoManagementList(GetVideoManagementListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DeleteVideoByID 删除视频根据id
func (C Controllers) DeleteVideoByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	DeleteVideoByIDReceive := new(receive.DeleteVideoByIDReceiveStruct)
	if err := ctx.ShouldBind(DeleteVideoByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.DeleteVideoByID(DeleteVideoByIDReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//LikeVideo 给视频点赞
func (C Controllers) LikeVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	LikeVideoReceive := new(receive.LikeVideoReceiveStruct)
	if err := ctx.ShouldBind(LikeVideoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.LikeVideo(LikeVideoReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
