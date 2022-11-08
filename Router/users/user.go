package users

import (
	"Go-Live/Controllers/users"
	"Go-Live/Middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(Middlewares.VerificationToken())
	{
		usersControllers := new(users.UserControllers)
		router.POST("/getUserInfo", usersControllers.GetUserInfo)

	}
}
