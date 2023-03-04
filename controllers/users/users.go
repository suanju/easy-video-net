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
	uid := ctx.GetUint("uid")
	results, err := users.GetUserInfo(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SetUserInfo  设置用户信息
func (us UserControllers) SetUserInfo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	setUserInfoReceive := new(receive.SetUserInfoReceiveStruct)
	if err := ctx.ShouldBind(setUserInfoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SetUserInfo(setUserInfoReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DetermineNameExists 判断名字是否存在
func (us UserControllers) DetermineNameExists(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	determineNameExistsReceive := new(receive.DetermineNameExistsStruct)
	if err := ctx.ShouldBind(determineNameExistsReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.DetermineNameExists(determineNameExistsReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Upload  文件上传
func (us UserControllers) Upload(ctx *gin.Context) {

	uid := ctx.GetUint("uid")
	file, _ := ctx.FormFile("file")
	results, err := users.Upload(file, uid, ctx)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//UpdateAvatar 修改头像
func (us UserControllers) UpdateAvatar(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	updateAvatarReceive := new(receive.UpdateAvatarStruct)
	if err := ctx.ShouldBind(updateAvatarReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.UpdateAvatar(updateAvatarReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetLiveData 获取直播资料
func (us UserControllers) GetLiveData(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetLiveData(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SaveLiveData 修改直播资料
func (us UserControllers) SaveLiveData(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	saveLiveDataReceive := new(receive.SaveLiveDataReceiveStruct)
	if err := ctx.ShouldBind(saveLiveDataReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.SaveLiveData(saveLiveDataReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//SendEmailVerificationCodeByChangePassword 找回密码发送验证码
func (us UserControllers) SendEmailVerificationCodeByChangePassword(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.SendEmailVerificationCodeByChangePassword(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//ChangePassword 修改密码
func (us UserControllers) ChangePassword(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	ChangePasswordReceive := new(receive.ChangePasswordReceiveStruct)
	if err := ctx.ShouldBind(ChangePasswordReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.ChangePassword(ChangePasswordReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//Attention 关注用户
func (us UserControllers) Attention(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	AttentionReceive := new(receive.AttentionReceiveStruct)
	if err := ctx.ShouldBind(AttentionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.Attention(AttentionReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//CreateFavorites 创建收藏夹
func (us UserControllers) CreateFavorites(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	CreateFavoritesReceive := new(receive.CreateFavoritesReceiveStruct)
	if err := ctx.ShouldBind(CreateFavoritesReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.CreateFavorites(CreateFavoritesReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DeleteFavorites 获取收藏夹
func (us UserControllers) DeleteFavorites(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	DeleteFavoritesReceive := new(receive.DeleteFavoritesReceiveStruct)
	if err := ctx.ShouldBind(DeleteFavoritesReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.DeleteFavorites(DeleteFavoritesReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetFavoritesList 获取收藏夹列表
func (us UserControllers) GetFavoritesList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetFavoritesList(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//FavoriteVideo 收藏视频
func (us UserControllers) FavoriteVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	FavoriteVideoReceive := new(receive.FavoriteVideoReceiveStruct)
	if err := ctx.ShouldBind(FavoriteVideoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.FavoriteVideo(FavoriteVideoReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetFavoritesListByFavoriteVideo 获取收藏夹列表在视频页面
func (us UserControllers) GetFavoritesListByFavoriteVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetFavoritesListByFavoriteVideoReceive := new(receive.GetFavoritesListByFavoriteVideoReceiveStruct)
	if err := ctx.ShouldBind(GetFavoritesListByFavoriteVideoReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.GetFavoritesListByFavoriteVideo(GetFavoritesListByFavoriteVideoReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetFavoriteVideoList 获取收藏夹视频列表
func (us UserControllers) GetFavoriteVideoList(ctx *gin.Context) {
	GetFavoriteVideoListReceive := new(receive.GetFavoriteVideoListReceiveStruct)
	if err := ctx.ShouldBind(GetFavoriteVideoListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.GetFavoriteVideoList(GetFavoriteVideoListReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetRecordList 获取历史记录
func (us UserControllers) GetRecordList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	GetRecordListReceive := new(receive.GetRecordListReceiveStruct)
	if err := ctx.ShouldBind(GetRecordListReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.GetRecordList(GetRecordListReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//ClearRecord 清空历史记录
func (us UserControllers) ClearRecord(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.ClearRecord(uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//DeleteRecordByID 删除历史记录根据id
func (us UserControllers) DeleteRecordByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	DeleteRecordByIDReceive := new(receive.DeleteRecordByIDReceiveStruct)
	if err := ctx.ShouldBind(DeleteRecordByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := users.DeleteRecordByID(DeleteRecordByIDReceive, uid)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
