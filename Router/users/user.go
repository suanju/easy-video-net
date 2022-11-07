package users

import (
	"Go-Live/Controllers/users"
	"Go-Live/Middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	_ = Router.Group("user").Use(Middlewares.VerificationToken())
	{
		_ = new(users.UserControllers)
	}
}
