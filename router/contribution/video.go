package contribution

import (
	"Go-Live/controllers/contribution"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	contributionControllers := new(contribution.Controllers)
	//不需要登入
	contributionRouterNoVerification := Router.Group("contribution").Use()
	{

		contributionRouterNoVerification.POST("/getVideoContributionByID", contributionControllers.GetVideoContributionByID)
		contributionRouterNoVerification.GET("/video/barrage/v3/", contributionControllers.GetVideoBarrage)
		contributionRouterNoVerification.GET("/getVideoBarrageList", contributionControllers.GetVideoBarrageList)
		contributionRouterNoVerification.POST("/getVideoComment", contributionControllers.GetVideoComment)
	}
	//需要登入 参数携带
	contributionRouterParameter := Router.Group("contribution").Use(middlewares.VerificationTokenAsParameter())
	{
		contributionRouterParameter.POST("/video/barrage/v3/", contributionControllers.SendVideoBarrage)
	}
	//请求头携带
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionRouter.POST("/createVideoContribution", contributionControllers.CreateVideoContribution)
		contributionRouter.POST("/videoPostComment", contributionControllers.VideoPostComment)
	}

}
