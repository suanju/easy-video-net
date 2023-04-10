package router

import (
	"easy-video-net/middlewares"
	"easy-video-net/router/callback"
	"easy-video-net/router/commonality"
	"easy-video-net/router/contribution"
	"easy-video-net/router/home"
	"easy-video-net/router/live"
	usersRouter "easy-video-net/router/users"
	"easy-video-net/router/ws"
	"github.com/gin-gonic/gin"
)

type RoutersGroup struct {
	Users        usersRouter.RouterGroup
	Live         live.RouterGroup
	Home         home.RouterGroup
	Commonality  commonality.RouterGroup
	Contribution contribution.RouterGroup
	Ws           ws.RouterGroup
	Callback     callback.RouterGroup
}

var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	router.Use(middlewares.Cors())
	PrivateGroup := router.Group("")
	PrivateGroup.Use()
	{
		//静态资源访问
		router.Static("/assets", "./assets/")
		RoutersGroupApp.Users.LoginRouter.InitLoginRouter(PrivateGroup)
		RoutersGroupApp.Users.SpaceRouter.InitSpaceRouter(PrivateGroup)
		RoutersGroupApp.Ws.InitSocketRouter(PrivateGroup)
		RoutersGroupApp.Users.InitRouter(PrivateGroup)
		RoutersGroupApp.Live.InitLiveRouter(PrivateGroup)
		RoutersGroupApp.Home.InitLiveRouter(PrivateGroup)
		RoutersGroupApp.Commonality.InitRouter(PrivateGroup)
		RoutersGroupApp.Contribution.VideoRouter.InitVideoRouter(PrivateGroup)
		RoutersGroupApp.Contribution.ArticleRouter.InitArticleRouter(PrivateGroup)
		RoutersGroupApp.Contribution.DiscussRouter.InitDiscussRouter(PrivateGroup)
		RoutersGroupApp.Callback.InitRouter(PrivateGroup)
	}

	err := router.Run()
	if err != nil {
		return
	}
}
