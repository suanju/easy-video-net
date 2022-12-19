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
		contributionRouter.POST("/GetArticleContributionListByUser", contributionControllers.GetArticleContributionListByUser)
		contributionRouter.POST("/GetArticleContributionByID", contributionControllers.GetArticleContributionByID)
	}
}
