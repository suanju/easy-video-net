package users

import (
	"Go-Live/controllers/users"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type SpaceRouter struct {
}

func (s *SpaceRouter) InitSpaceRouter(Router *gin.RouterGroup) {
	//非必须登入
	spaceRouter := Router.Group("space").Use(middlewares.VerificationTokenNotNecessary())
	{
		spaceControllers := new(users.SpaceControllers)
		spaceRouter.POST("/getSpaceIndividual", spaceControllers.GetSpaceIndividual)
		spaceRouter.POST("/getReleaseInformation", spaceControllers.GetReleaseInformation)
	}
}
