package contribution

import (
	"easy-video-net/controllers/contribution"
	"easy-video-net/middlewares"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	contributionControllers := new(contribution.Controllers)
	//不需要登入
	contributionRouterNoVerification := Router.Group("contribution").Use()
	{
		contributionRouterNoVerification.GET("/video/barrage/v3/", contributionControllers.GetVideoBarrage)
		contributionRouterNoVerification.GET("/getVideoBarrageList", contributionControllers.GetVideoBarrageList)
		contributionRouterNoVerification.POST("/getVideoComment", contributionControllers.GetVideoComment)
	}
	//非必须登入
	contributionRouterNotNecessary := Router.Group("contribution").Use(middlewares.VerificationTokenNotNecessary())
	{
		contributionRouterNotNecessary.POST("/getVideoContributionByID", contributionControllers.GetVideoContributionByID)
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
		contributionRouter.POST("/updateVideoContribution", contributionControllers.UpdateVideoContribution)
		contributionRouter.POST("/deleteVideoByID", contributionControllers.DeleteVideoByID)
		contributionRouter.POST("/videoPostComment", contributionControllers.VideoPostComment)
		contributionRouter.POST("/getVideoManagementList", contributionControllers.GetVideoManagementList)
		contributionRouter.POST("/likeVideo", contributionControllers.LikeVideo)
	}

}
