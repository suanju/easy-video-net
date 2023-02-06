package middlewares

import (
	"Go-Live/models/users"
	"Go-Live/utils/jwt"
	ControllersCommon "Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//VerificationToken 请求头中携带token
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
		if !u.IsExistByField("id", claim.UserID) {
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

//VerificationTokenAsParameter
func VerificationTokenAsParameter() gin.HandlerFunc {
	type qu struct {
		Token string `json:"token"`
	}
	return func(c *gin.Context) {
		req := new(qu)
		if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
			validator.CheckParams(c, err)
			return
		}
		token := req.Token
		claim, err := jwt.ParseToken(token)
		if err != nil {
			ControllersCommon.NotLogin(c, "Token过期")
			c.Abort()
			return
		}
		u := new(users.User)
		if !u.IsExistByField("id", claim.UserID) {
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
