package users

import (
	"Go-Live/controllers/users"
	"Go-Live/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middlewares.VerificationToken())
	{
		usersControllers := new(users.UserControllers)
		router.POST("/getUserInfo", usersControllers.GetUserInfo)
		router.POST("/setUserInfo", usersControllers.SetUserInfo)
		router.POST("/determineNameExists", usersControllers.DetermineNameExists)
		router.POST("/upload", usersControllers.Upload)

	}
}
