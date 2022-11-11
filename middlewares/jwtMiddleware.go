package middlewares

import (
	"Go-Live/models/users"
	"Go-Live/utils/jwt"
	ControllersCommon "Go-Live/utils/response"
	"github.com/gin-gonic/gin"
)

func VerificationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ControllersCommon.NotLogin(c, "Token过期")
			c.Abort()
			return
		}
		u := new(users.User)
		if !u.IsExistByField("id", claim.UserId) {
			//没有改用户的情况下
			ControllersCommon.NotLogin(c, "用户异常")
			c.Abort()
			return
		}
		c.Set("currentUserID", u.ID)
		c.Set("currentUserName", u.Username)
		c.Next()
	}
}
