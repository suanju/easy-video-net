package contribution

import (
	"easy-video-net/controllers/contribution"
	"easy-video-net/middlewares"
	"github.com/gin-gonic/gin"
)

type DiscussRouter struct {
}

func (v *DiscussRouter) InitDiscussRouter(Router *gin.RouterGroup) {
	contributionControllers := new(contribution.Controllers)
	//请求头携带
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionRouter.POST("/getDiscussVideoList", contributionControllers.GetDiscussVideoList)
		contributionRouter.POST("/getDiscussArticleList", contributionControllers.GetDiscussArticleList)
		contributionRouter.POST("/getDiscussBarrageList", contributionControllers.GetDiscussBarrageList)
	}

}
