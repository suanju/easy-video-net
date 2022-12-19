package contribution

import (
	"Go-Live/logic/contribution"
	"Go-Live/models/contribution/article"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

//CreateArticleContribution 发布专栏
func (C Controllers) CreateArticleContribution(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	CreateArticleContributionReceive := new(article.CreateArticleContributionReceiveStruct)
	if err := ctx.ShouldBind(CreateArticleContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.CreateArticleContribution(CreateArticleContributionReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleContributionListByUser 查询用户发布的专栏
func (C Controllers) GetArticleContributionListByUser(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	GetArticleContributionListByUserReceive := new(article.GetArticleContributionListByUserReceiveStruct)
	if err := ctx.ShouldBind(GetArticleContributionListByUserReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleContributionListByUser(GetArticleContributionListByUserReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleContributionByID 查询专栏信息根据ID
func (C Controllers) GetArticleContributionByID(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	GetArticleContributionByIDReceive := new(article.GetArticleContributionByIDReceiveStruct)
	if err := ctx.ShouldBind(GetArticleContributionByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleContributionByID(GetArticleContributionByIDReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
