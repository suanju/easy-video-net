package users

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/users"
	"Go-Live/logic/users"
	"github.com/gin-gonic/gin"
)

type SpaceControllers struct {
	controllers.BaseControllers
}

//GetSpaceIndividual 获取个人空间
func (sp SpaceControllers) GetSpaceIndividual(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetSpaceIndividualReceiveStruct)); err == nil {
		results, err := users.GetSpaceIndividual(rec, uid)
		sp.Response(ctx, results, err)
	}
}

//GetReleaseInformation 获取发布信息（视频and专栏）
func (sp SpaceControllers) GetReleaseInformation(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetReleaseInformationReceiveStruct)); err == nil {
		results, err := users.GetReleaseInformation(rec)
		sp.Response(ctx, results, err)
	}
}

//GetAttentionList 获取关注列表
func (sp SpaceControllers) GetAttentionList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetAttentionListReceiveStruct)); err == nil {
		results, err := users.GetAttentionList(rec, uid)
		sp.Response(ctx, results, err)
	}
}

//GetVermicelliList 获取粉丝列表
func (sp SpaceControllers) GetVermicelliList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetVermicelliListReceiveStruct)); err == nil {
		results, err := users.GetVermicelliList(rec, uid)
		sp.Response(ctx, results, err)
	}
}
