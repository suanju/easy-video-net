package users

import (
	receive "Go-Live/interaction/receive/users"
	"Go-Live/logic/users"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
}

//GetUserInfo 获取用户信息
func (us UserControllers) GetUserInfo(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	results, err := users.GetUserInfo(userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SetUserInfo  设置用户信息
func (us UserControllers) SetUserInfo(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	setUserInfoReceive := new(receive.SetUserInfoReceiveStruct)
	if err := ctx.ShouldBind(setUserInfoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SetUserInfo(setUserInfoReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DetermineNameExists 判断名字是否存在
func (us UserControllers) DetermineNameExists(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	determineNameExistsReceive := new(receive.DetermineNameExistsStruct)
	if err := ctx.ShouldBind(determineNameExistsReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.DetermineNameExists(determineNameExistsReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Upload  文件上传
func (us UserControllers) Upload(ctx *gin.Context) {

	userID := ctx.GetUint("currentUserID")
	file, _ := ctx.FormFile("file")
	results, err := users.Upload(file, userID, ctx)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//UpdateAvatar 修改头像
func (us UserControllers) UpdateAvatar(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	updateAvatarReceive := new(receive.UpdateAvatarStruct)
	if err := ctx.ShouldBind(updateAvatarReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.UpdateAvatar(updateAvatarReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetLiveData 获取直播资料
func (us UserControllers) GetLiveData(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	results, err := users.GetLiveData(userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SaveLiveData 修改直播资料
func (us UserControllers) SaveLiveData(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	saveLiveDataReceive := new(receive.SaveLiveDataReceiveStruct)
	if err := ctx.ShouldBind(saveLiveDataReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SaveLiveData(saveLiveDataReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerificationCodeByChangePassword 找回密码发送验证码
func (us UserControllers) SendEmailVerificationCodeByChangePassword(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	results, err := users.SendEmailVerificationCodeByChangePassword(userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//ChangePassword 修改密码
func (us UserControllers) ChangePassword(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	ChangePasswordReceive := new(receive.ChangePasswordReceiveStruct)
	if err := ctx.ShouldBind(ChangePasswordReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.ChangePassword(ChangePasswordReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Attention 关注用户
func (us UserControllers) Attention(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	AttentionReceive := new(receive.AttentionReceiveStruct)
	if err := ctx.ShouldBind(AttentionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.Attention(AttentionReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
