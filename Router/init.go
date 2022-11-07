package Router

import (
	"Go-Live/Middlewares"
	"Go-Live/Router/live"
	usersRouter "Go-Live/Router/users"
	"github.com/gin-gonic/gin"
)

type RoutersGroup struct {
	Users usersRouter.RouterGroup
	Live  live.RouterGroup
}

var RoutersGroupApp = new(RoutersGroup)

func InitRouter() {
	router := gin.Default()
	router.Use(Middlewares.Cors())
	PrivateGroup := router.Group("")
	PrivateGroup.Use()
	{
		//静态资源访问
		router.Static("/assets", "./Assets/")
		RoutersGroupApp.Users.LoginRouter.InitLoginRouter(PrivateGroup)
		RoutersGroupApp.Users.InitRouter(PrivateGroup)
		RoutersGroupApp.Live.InitLiveRouter(PrivateGroup)
	}

	err := router.Run()
	if err != nil {
		return
	}
}
