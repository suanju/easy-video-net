package users

import (
	"easy-video-net/controllers"
	receive "easy-video-net/interaction/receive/users"
	"easy-video-net/logic/users"
	"github.com/gin-gonic/gin"
)

type LoginControllers struct {
	controllers.BaseControllers
}

//WxAuthorization 微信快捷登入
func (lg LoginControllers) WxAuthorization(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.WxAuthorizationReceiveStruct)); err == nil {
		results, err := users.WxAuthorization(rec)
		lg.Response(ctx, results, err)
	}
}

//Login 登入
func (lg LoginControllers) Login(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.LoginReceiveStruct)); err == nil {
		results, err := users.Login(rec)
		lg.Response(ctx, results, err)
	}
}

//Register 注册
func (lg LoginControllers) Register(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.RegisterReceiveStruct)); err == nil {
		results, err := users.Register(rec)
		lg.Response(ctx, results, err)
	}
}

//SendEmailVerCode 获取验证码(注册)
func (lg LoginControllers) SendEmailVerCode(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.SendEmailVerCodeReceiveStruct)); err == nil {
		results, err := users.SendEmailVerCode(rec)
		lg.Response(ctx, results, err)
	}
}

//SendEmailVerCodeByForget 获取邮箱验证码(忘记密码)
func (lg LoginControllers) SendEmailVerCodeByForget(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.SendEmailVerCodeReceiveStruct)); err == nil {
		results, err := users.SendEmailVerCodeByForget(rec)
		lg.Response(ctx, results, err)
	}
}

//Forget 找回密码
func (lg LoginControllers) Forget(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.ForgetReceiveStruct)); err == nil {
		results, err := users.Forget(rec)
		lg.Response(ctx, results, err)
	}
}
