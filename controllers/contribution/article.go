package contribution

import (
	receive "Go-Live/interaction/receive/contribution/article"
	"Go-Live/logic/contribution"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"

	"github.com/gin-gonic/gin"
)

//CreateArticleContribution 发布专栏
func (C Controllers) CreateArticleContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	CreateArticleContributionReceive := new(receive.CreateArticleContributionReceiveStruct)
	if err := ctx.ShouldBind(CreateArticleContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.CreateArticleContribution(CreateArticleContributionReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//UpdateArticleContribution 更新专栏
func (C Controllers) UpdateArticleContribution(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	UpdateArticleContributionReceive := new(receive.UpdateArticleContributionReceiveStruct)
	if err := ctx.ShouldBind(UpdateArticleContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.UpdateArticleContribution(UpdateArticleContributionReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DeleteArticleByID 删除专栏
func (C Controllers) DeleteArticleByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	DeleteArticleByIDReceive := new(receive.DeleteArticleByIDReceiveStruct)
	if err := ctx.ShouldBind(DeleteArticleByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.DeleteArticleByID(DeleteArticleByIDReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleContributionList 首页查询专栏
func (C Controllers) GetArticleContributionList(ctx *gin.Context) {
	GetArticleContributionListReceive := new(receive.GetArticleContributionListReceiveStruct)
	if err := ctx.ShouldBind(GetArticleContributionListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleContributionList(GetArticleContributionListReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleContributionListByUser 查询用户发布的专栏
func (C Controllers) GetArticleContributionListByUser(ctx *gin.Context) {
	GetArticleContributionListByUserReceive := new(receive.GetArticleContributionListByUserReceiveStruct)
	if err := ctx.ShouldBind(GetArticleContributionListByUserReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleContributionListByUser(GetArticleContributionListByUserReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleContributionByID 查询专栏信息根据ID
func (C Controllers) GetArticleContributionByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetArticleContributionByIDReceive := new(receive.GetArticleContributionByIDReceiveStruct)
	if err := ctx.ShouldBind(GetArticleContributionByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleContributionByID(GetArticleContributionByIDReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//ArticlePostComment 发布评论
func (C Controllers) ArticlePostComment(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	ArticlePostCommentReceive := new(receive.ArticlesPostCommentReceiveStruct)
	if err := ctx.ShouldBind(ArticlePostCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.ArticlePostComment(ArticlePostCommentReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleComment 获取文章评论
func (C Controllers) GetArticleComment(ctx *gin.Context) {
	GetArticleCommentReceive := new(receive.GetArticleCommentReceiveStruct)
	if err := ctx.ShouldBind(GetArticleCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleComment(GetArticleCommentReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleClassificationList 获取视频分类
func (C Controllers) GetArticleClassificationList(ctx *gin.Context) {
	results, err := contribution.GetArticleClassificationList()
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleTotalInfo 获取文章相关总和信息
func (C Controllers) GetArticleTotalInfo(ctx *gin.Context) {
	results, err := contribution.GetArticleTotalInfo()
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetArticleManagementList 创作中心获取专栏稿件列表
func (C Controllers) GetArticleManagementList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetArticleManagementListReceive := new(receive.GetArticleManagementListReceiveStruct)
	if err := ctx.ShouldBind(GetArticleManagementListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetArticleManagementList(GetArticleManagementListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
