package contribution

import (
	"easy-video-net/controllers"
	receive "easy-video-net/interaction/receive/contribution/discuss"
	"easy-video-net/logic/contribution"
	"github.com/gin-gonic/gin"
)

//GetDiscussVideoList 获取视频评论列表
func (c Controllers) GetDiscussVideoList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetDiscussVideoListReceiveStruct)); err == nil {
		results, err := contribution.GetDiscussVideoList(rec, uid)
		c.Response(ctx, results, err)
	}
}

//GetDiscussArticleList 获取专栏评论列表
func (c Controllers) GetDiscussArticleList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetDiscussArticleListReceiveStruct)); err == nil {
		results, err := contribution.GetDiscussArticleList(rec, uid)
		c.Response(ctx, results, err)
	}
}

//GetDiscussBarrageList 获取视频弹幕列表
func (c Controllers) GetDiscussBarrageList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetDiscussBarrageListReceiveStruct)); err == nil {
		results, err := contribution.GetDiscussBarrageList(rec, uid)
		c.Response(ctx, results, err)
	}
}
