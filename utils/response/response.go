package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
	Version interface{} `json:"version,omitempty"`
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

func Default(ctx *gin.Context) {
	rd := &Data{
		Code:    CodeDefault,
		Message: CodeDefault.Msg(),
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

//BarrageSuccess 弹幕播放器响应
func BarrageSuccess(ctx *gin.Context, data interface{}) {
	rd := &Data{
		Code:    0,
		Message: CodeSuccess.Msg(),
		Data:    data,
		Version: 3,
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
