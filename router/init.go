package router

import (
	"Go-Live/middlewares"
	"Go-Live/router/commonality"
	"Go-Live/router/contribution"
	"Go-Live/router/home"
	"Go-Live/router/live"
	usersRouter "Go-Live/router/users"
	"Go-Live/router/ws"
	"github.com/gin-gonic/gin"
)

type RoutersGroup struct {
	Users        usersRouter.RouterGroup
	Live         live.RouterGroup
	Home         home.RouterGroup
	Commonality  commonality.RouterGroup
	Contribution contribution.RouterGroup
	Ws           ws.Router
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
	}

	err := router.Run()
	if err != nil {
		return
	}
}
