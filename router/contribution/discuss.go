package contribution

import (
	"Go-Live/controllers/contribution"
	"Go-Live/middlewares"
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
	}

}
