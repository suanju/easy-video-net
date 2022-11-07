package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func Error(ctx *gin.Context, msg string) {
	rd := &Data{
		Code:    CodeServerBusy,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}
func TypeError(ctx *gin.Context, msg string) {
	rd := &Data{
		Code:    CodeTypeError,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func NotLogin(ctx *gin.Context, msg string) {
	rd := &Data{
		Code:    CodeNotLogin,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &Data{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func Success(ctx *gin.Context, data interface{}) {
	rd := &Data{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
