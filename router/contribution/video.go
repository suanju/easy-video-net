package contribution

import (
	"Go-Live/controllers/contribution"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	contributionRouter := Router.Group("contribution").Use(middlewares.VerificationToken())
	{
		contributionControllers := new(contribution.Controllers)
		contributionRouter.POST("/createVideoContribution", contributionControllers.CreateVideoContribution)
		contributionRouter.POST("/getVideoContributionByID", contributionControllers.GetVideoContributionByID)
	}
}
