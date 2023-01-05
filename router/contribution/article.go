package contribution

import (
	"Go-Live/controllers/contribution"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (v *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionControllers := new(contribution.Controllers)
		contributionRouter.POST("/createArticleContribution", contributionControllers.CreateArticleContribution)
		contributionRouter.POST("/getArticleContributionListByUser", contributionControllers.GetArticleContributionListByUser)
		contributionRouter.POST("/getArticleContributionByID", contributionControllers.GetArticleContributionByID)
		contributionRouter.POST("/articlePostComment", contributionControllers.ArticlePostComment)
		contributionRouter.POST("/getArticleComment", contributionControllers.GetArticleComment)

	}
}
