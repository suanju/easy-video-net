package live

import (
	"easy-video-net/controllers/live"
	"easy-video-net/middlewares"
	"github.com/gin-gonic/gin"
)

type LivesRouter struct {
}

func (s *LivesRouter) InitLiveRouter(Router *gin.RouterGroup) {
	liveRouter := Router.Group("live").Use(middlewares.VerificationToken())
	{
		liveControllers := new(live.LivesControllers)
		liveRouter.POST("/getLiveRoom", liveControllers.GetLiveRoom)
		liveRouter.POST("/getLiveRoomInfo", liveControllers.GetLiveRoomInfo)
		liveRouter.POST("/getBeLiveList", liveControllers.GetBeLiveList)
	}
}
