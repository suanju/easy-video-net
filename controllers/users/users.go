package users

import (
	"Go-Live/logic/users"
	usersModel "Go-Live/models/users"
	"Go-Live/models/users/liveInfo"
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
	setUserInfoReceive := new(usersModel.SetUserInfoStruct)
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
	determineNameExistsReceive := new(usersModel.DetermineNameExistsStruct)
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
	updateAvatarReceive := new(usersModel.UpdateAvatarStruct)
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
	saveLiveDataReceive := new(liveInfo.SaveLiveDataStruct)
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
