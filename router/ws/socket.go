package ws

import (
	"Go-Live/controllers/contribution"
	"Go-Live/controllers/live"
	"Go-Live/controllers/users"
	"Go-Live/middlewares"
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
