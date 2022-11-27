package live

import (
	"Go-Live/controllers/live"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type LivesRouter struct {
}

func (s *LivesRouter) InitLiveRouter(Router *gin.RouterGroup) {
	liveRouter := Router.Group("live").Use(middlewares.VerificationToken())
	{
		liveControllers := new(live.LivesControllers)
		liveRouter.POST("/getLiveRoom", liveControllers.GetLiveRoom)
		liveRouter.POST("/liveSocket", liveControllers.LiveSocket)
	}
}
