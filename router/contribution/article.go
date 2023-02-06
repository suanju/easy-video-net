package contribution

import (
	"Go-Live/controllers/contribution"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (v *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	contributionControllers := new(contribution.Controllers)
	contributionRouterNoVerification := Router.Group("contribution").Use()
	{
		contributionRouterNoVerification.POST("/getArticleContributionListByUser", contributionControllers.GetArticleContributionListByUser)
		contributionRouterNoVerification.POST("/getArticleContributionByID", contributionControllers.GetArticleContributionByID)
		contributionRouterNoVerification.POST("/getArticleComment", contributionControllers.GetArticleComment)
		contributionRouterNoVerification.POST("/getArticleClassificationList", contributionControllers.GetArticleClassificationList)
		contributionRouterNoVerification.POST("/getArticleTotalInfo", contributionControllers.GetArticleTotalInfo)

	}
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionRouter.POST("/createArticleContribution", contributionControllers.CreateArticleContribution)
		contributionRouter.POST("/articlePostComment", contributionControllers.ArticlePostComment)
	}
}
