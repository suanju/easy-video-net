package users

import (
	"easy-video-net/controllers/users"
	"easy-video-net/middlewares"

	"github.com/gin-gonic/gin"
)

type SpaceRouter struct {
}

func (s *SpaceRouter) InitSpaceRouter(Router *gin.RouterGroup) {

	spaceControllers := new(users.SpaceControllers)

	//必须登入
	spaceRouter := Router.Group("space").Use(middlewares.VerificationToken())
	{
		spaceRouter.POST("/getAttentionList", spaceControllers.GetAttentionList)
		spaceRouter.POST("/getVermicelliList", spaceControllers.GetVermicelliList)
	}

	//非必须登入
	spaceRouterNotNecessary := Router.Group("space").Use(middlewares.VerificationTokenNotNecessary())
	{
		spaceRouterNotNecessary.POST("/getSpaceIndividual", spaceControllers.GetSpaceIndividual)
		spaceRouterNotNecessary.POST("/getReleaseInformation", spaceControllers.GetReleaseInformation)
	}
}
