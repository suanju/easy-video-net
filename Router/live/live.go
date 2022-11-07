package live

import (
	"Go-Live/Controllers/live"
	"Go-Live/Middlewares"
	"github.com/gin-gonic/gin"
)

type LivesRouter struct {
}

func (s *LivesRouter) InitLiveRouter(Router *gin.RouterGroup) {
	liveRouter := Router.Group("live").Use(Middlewares.VerificationToken())
	{
		liveControllers := new(live.LivesControllers)
		liveRouter.POST("/getLiveRoom", liveControllers.GetLiveRoom)
	}
}
