package middlewares

import (
	"Go-Live/models/users"
	"Go-Live/utils/jwt"
	"Go-Live/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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
			response.NotLoginWs(conn, "Token 验证失败")
			_ = conn.Close()
			c.Abort()
			return
		}
		u := new(users.User)
		if !u.IsExistByField("id", claim.UserID) {
			//没有改用户的情况下
			response.NotLoginWs(conn, "用户异常")
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
