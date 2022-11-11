package users

import (
	"Go-Live/controllers"
	"Go-Live/logic/users"
	usersModel "Go-Live/models/users"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type LoginControllers struct {
	controllers.BaseControllers
}

//WxAuthorization 微信快捷登入
func (lg LoginControllers) WxAuthorization(ctx *gin.Context) {
	WxAuthorizationReceive := new(usersModel.WxAuthorizationReceiveStruct)
	if err := ctx.ShouldBind(WxAuthorizationReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.WxAuthorization(WxAuthorizationReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Login 登入
func (lg LoginControllers) Login(ctx *gin.Context) {
	LoginReceive := new(usersModel.LoginReceiveStruct)
	if err := ctx.ShouldBind(LoginReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.Login(LoginReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Register 注册
func (lg LoginControllers) Register(ctx *gin.Context) {
	RegisterReceive := new(usersModel.RegisterReceiveStruct)
	if err := ctx.ShouldBind(RegisterReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.Register(RegisterReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerCode 获取验证码(注册)
func (lg LoginControllers) SendEmailVerCode(ctx *gin.Context) {
	SendEmailVerCodeReceive := new(usersModel.SendEmailVerCodeReceiveStruct)
	if err := ctx.ShouldBind(SendEmailVerCodeReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SendEmailVerCode(SendEmailVerCodeReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerCodeByForget 获取邮箱验证码(忘记密码)
func (lg LoginControllers) SendEmailVerCodeByForget(ctx *gin.Context) {
	SendEmailVerCodeReceive := new(usersModel.SendEmailVerCodeReceiveStruct)
	if err := ctx.ShouldBind(SendEmailVerCodeReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SendEmailVerCodeByForget(SendEmailVerCodeReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Forget 找回密码
func (lg LoginControllers) Forget(ctx *gin.Context) {

	ForgetReceive := new(usersModel.ForgetReceiveStruct)
	if err := ctx.ShouldBind(ForgetReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.Forget(ForgetReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
