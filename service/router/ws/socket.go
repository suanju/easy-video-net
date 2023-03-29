package ws

import (
	"easy-video-net/controllers/contribution"
	"easy-video-net/controllers/live"
	"easy-video-net/controllers/users"
	"easy-video-net/middlewares"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) InitSocketRouter(Router *gin.RouterGroup) {
	socketRouter := Router.Group("ws").Use(middlewares.VerificationTokenAsSocket())
	{
		userControllers := new(users.UserControllers)
		liveControllers := new(live.LivesControllers)
		contributionControllers := new(contribution.Controllers)
		socketRouter.GET("/noticeSocket", userControllers.NoticeSocket)
		socketRouter.GET("/chatSocket", userControllers.ChatSocket)
		socketRouter.GET("/chatByUserSocket", userControllers.ChatByUserSocket)
		socketRouter.GET("/liveSocket", liveControllers.LiveSocket)
		socketRouter.GET("/videoSocket", contributionControllers.VideoSocket)
	}
}
