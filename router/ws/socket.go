package ws

import (
	"Go-Live/controllers/live"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) InitSocketRouter(Router *gin.RouterGroup) {
	liveRouter := Router.Group("ws").Use(middlewares.VerificationTokenAsSocket())
	{
		liveControllers := new(live.LivesControllers)
		liveRouter.GET("/liveSocket", liveControllers.LiveSocket)
	}
}
