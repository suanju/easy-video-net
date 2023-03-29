package middlewares

import (
	"easy-video-net/models/users"
	"easy-video-net/utils/jwt"
	ControllersCommon "easy-video-net/utils/response"
	Response "easy-video-net/utils/response"
	"easy-video-net/utils/validator"
	"github.com/gorilla/websocket"
	"net/http"

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
		c.Set("uid", u.ID)
		c.Set("currentUserName", u.Username)
		c.Next()
	}
}

//VerificationTokenAsParameter body参数中携带token
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
		c.Set("uid", u.ID)
		c.Set("currentUserName", u.Username)
		c.Next()
	}
}

//VerificationTokenNotNecessary 请求头中携带token (非必须)
func VerificationTokenNotNecessary() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			//用户未登入时不验证
			c.Next()
		} else {
			//用户登入情况
			claim, err := jwt.ParseToken(token)
			if err != nil {
				c.Next()
			}
			u := new(users.User)
			if !u.IsExistByField("id", claim.UserID) {
				//没有改用户的情况下
				ControllersCommon.NotLogin(c, "用户异常")
				c.Abort()
				return
			}
			c.Set("uid", u.ID)
			c.Set("currentUserName", u.Username)
			c.Next()
		}
	}
}

func VerificationTokenAsSocket() gin.HandlerFunc {
	return func(c *gin.Context) {
		//升级ws 以便返回消息
		conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			http.NotFound(c.Writer, c.Request)
			c.Abort()
			return
		}
		token := c.Query("token")
		claim, err := jwt.ParseToken(token)
		if err != nil {
			Response.NotLoginWs(conn, "Token 验证失败")
			_ = conn.Close()
			c.Abort()
			return
		}
		u := new(users.User)
		if !u.IsExistByField("id", claim.UserID) {
			//没有改用户的情况下
			Response.NotLoginWs(conn, "用户异常")
			_ = conn.Close()
			c.Abort()
			return
		}
		c.Set("uid", u.ID)
		c.Set("conn", conn)
		c.Set("currentUserName", u.Username)
		c.Next()
	}
}
