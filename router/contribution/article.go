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
		contributionRouterNoVerification.POST("/createArticleContribution", contributionControllers.CreateArticleContribution)
		contributionRouterNoVerification.POST("/getArticleContributionListByUser", contributionControllers.GetArticleContributionListByUser)
		contributionRouterNoVerification.POST("/getArticleContributionByID", contributionControllers.GetArticleContributionByID)
		contributionRouterNoVerification.POST("/getArticleComment", contributionControllers.GetArticleComment)

	}
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionRouter.POST("/articlePostComment", contributionControllers.ArticlePostComment)
	}
}
