package home

import (
	"Go-Live/controllers/home"
	"github.com/gin-gonic/gin"
)

type homeRouter struct {
}

func (s *homeRouter) InitLiveRouter(Router *gin.RouterGroup) {
	homeRouter := Router.Group("home")
	{
		homeControllers := new(home.Controllers)
		homeRouter.POST("/getHomeInfo", homeControllers.GetHomeInfo)
	}
}
