package contribution

import (
	"easy-video-net/controllers"
	receive "easy-video-net/interaction/receive/contribution/video"
	"easy-video-net/logic/contribution"
	"easy-video-net/utils/response"
	"easy-video-net/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controllers struct {
	controllers.BaseControllers
}

//CreateVideoContribution 发布视频和编辑视频
func (c Controllers) CreateVideoContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.CreateVideoContributionReceiveStruct)); err == nil {
		results, err := contribution.CreateVideoContribution(rec, uid)
		c.Response(ctx, results, err)
	}
}

//UpdateVideoContribution 编辑视频
func (c Controllers) UpdateVideoContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.UpdateVideoContributionReceiveStruct)); err == nil {
		results, err := contribution.UpdateVideoContribution(rec, uid)
		c.Response(ctx, results, err)
	}
}

// GetVideoContributionByID  根据id获取视频信息
func (c Controllers) GetVideoContributionByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetVideoContributionByIDReceiveStruct)); err == nil {
		results, err := contribution.GetVideoContributionByID(rec, uid)
		c.Response(ctx, results, err)
	}

}

// SendVideoBarrage  发送视频弹幕
func (c Controllers) SendVideoBarrage(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	SendVideoBarrageReceive := new(receive.SendVideoBarrageReceiveStruct)
	//使用ShouldBindBodyWith解决重复bind问题
	if err := ctx.ShouldBindBodyWith(SendVideoBarrageReceive, binding.JSON); err == nil {
		results, err := contribution.SendVideoBarrage(SendVideoBarrageReceive, uid)
		if err != nil {
			response.Error(ctx, err.Error())
			return
		}
		response.BarrageSuccess(ctx, results)
	} else {
		validator.CheckParams(ctx, err)
	}
}

// GetVideoBarrage  获取视频弹幕 (播放器）
func (c Controllers) GetVideoBarrage(ctx *gin.Context) {
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
func (c Controllers) GetVideoBarrageList(ctx *gin.Context) {
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
func (c Controllers) VideoPostComment(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.VideosPostCommentReceiveStruct)); err == nil {
		results, err := contribution.VideoPostComment(rec, uid)
		c.Response(ctx, results, err)
	}
}

//GetVideoComment 获取视频评论
func (c Controllers) GetVideoComment(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetVideoCommentReceiveStruct)); err == nil {
		results, err := contribution.GetVideoComment(rec)
		c.Response(ctx, results, err)
	}
}

//GetVideoManagementList 创作中心获取视频稿件列表
func (c Controllers) GetVideoManagementList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetVideoManagementListReceiveStruct)); err == nil {
		results, err := contribution.GetVideoManagementList(rec, uid)
		c.Response(ctx, results, err)
	}
}

//DeleteVideoByID 删除视频根据id
func (c Controllers) DeleteVideoByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.DeleteVideoByIDReceiveStruct)); err == nil {
		results, err := contribution.DeleteVideoByID(rec, uid)
		c.Response(ctx, results, err)
	}
}

//LikeVideo 给视频点赞
func (c Controllers) LikeVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.LikeVideoReceiveStruct)); err == nil {
		results, err := contribution.LikeVideo(rec, uid)
		c.Response(ctx, results, err)
	}
}
