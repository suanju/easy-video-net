package users

import (
	"Go-Live/Controllers/users"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
}

func (s *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	usersRouter := Router.Group("login").Use()
	{
		loginControllers := new(users.LoginControllers)
		usersRouter.POST("/wxAuthorization", loginControllers.WxAuthorization)
		usersRouter.POST("/register", loginControllers.Register)
		usersRouter.POST("/login", loginControllers.Login)
		usersRouter.POST("/sendEmailVerificationCode", loginControllers.SendEmailVerCode)
		usersRouter.POST("/sendEmailVerificationCodeByForget", loginControllers.SendEmailVerCodeByForget)
		usersRouter.POST("/forget", loginControllers.Forget)
	}
}
