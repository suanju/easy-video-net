package router

import (
	"Go-Live/middlewares"
	"Go-Live/router/commonality"
	"Go-Live/router/live"
	usersRouter "Go-Live/router/users"
	"github.com/gin-gonic/gin"
)

type RoutersGroup struct {
	Users       usersRouter.RouterGroup
	Live        live.RouterGroup
	Commonality commonality.RouterGroup
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
		RoutersGroupApp.Users.InitRouter(PrivateGroup)
		RoutersGroupApp.Live.InitLiveRouter(PrivateGroup)
		RoutersGroupApp.Commonality.InitRouter(PrivateGroup)
	}

	err := router.Run()
	if err != nil {
		return
	}
}
