package contribution

import (
	receive "Go-Live/interaction/receive/contribution/discuss"
	"Go-Live/logic/contribution"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"

	"github.com/gin-gonic/gin"
)

//GetDiscussVideoList 获取视频评论列表
func (C Controllers) GetDiscussVideoList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetDiscussVideoListReceive := new(receive.GetDiscussVideoListReceiveStruct)
	if err := ctx.ShouldBind(GetDiscussVideoListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetDiscussVideoList(GetDiscussVideoListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetDiscussArticleList 获取专栏评论列表
func (C Controllers) GetDiscussArticleList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetDiscussArticleListReceive := new(receive.GetDiscussArticleListReceiveStruct)
	if err := ctx.ShouldBind(GetDiscussArticleListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetDiscussArticleList(GetDiscussArticleListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetDiscussBarrageList 获取视频弹幕列表
func (C Controllers) GetDiscussBarrageList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetDiscussBarrageListReceive := new(receive.GetDiscussBarrageListReceiveStruct)
	if err := ctx.ShouldBind(GetDiscussBarrageListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetDiscussBarrageList(GetDiscussBarrageListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
