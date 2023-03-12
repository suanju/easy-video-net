package users

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/users"
	"Go-Live/logic/users"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
	controllers.BaseControllers
}

//GetUserInfo 获取用户信息
func (us UserControllers) GetUserInfo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetUserInfo(uid)
	us.Response(ctx, results, err)
}

//SetUserInfo  设置用户信息
func (us UserControllers) SetUserInfo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.SetUserInfoReceiveStruct)); err == nil {
		results, err := users.SetUserInfo(rec, uid)
		us.Response(ctx, results, err)
	}
}

//DetermineNameExists 判断名字是否存在
func (us UserControllers) DetermineNameExists(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.DetermineNameExistsStruct)); err == nil {
		results, err := users.DetermineNameExists(rec, uid)
		us.Response(ctx, results, err)
	}
}

//Upload  文件上传
func (us UserControllers) Upload(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	file, _ := ctx.FormFile("file")
	results, err := users.Upload(file, uid, ctx)
	us.Response(ctx, results, err)
}

//UpdateAvatar 修改头像
func (us UserControllers) UpdateAvatar(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.UpdateAvatarStruct)); err == nil {
		results, err := users.UpdateAvatar(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetLiveData 获取直播资料
func (us UserControllers) GetLiveData(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetLiveData(uid)
	us.Response(ctx, results, err)
}

//SaveLiveData 修改直播资料
func (us UserControllers) SaveLiveData(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.SaveLiveDataReceiveStruct)); err == nil {
		results, err := users.SaveLiveData(rec, uid)
		us.Response(ctx, results, err)
	}
}

//SendEmailVerificationCodeByChangePassword 找回密码发送验证码
func (us UserControllers) SendEmailVerificationCodeByChangePassword(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.SendEmailVerificationCodeByChangePassword(uid)
	us.Response(ctx, results, err)
}

//ChangePassword 修改密码
func (us UserControllers) ChangePassword(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.ChangePasswordReceiveStruct)); err == nil {
		results, err := users.ChangePassword(rec, uid)
		us.Response(ctx, results, err)
	}
}

//Attention 关注用户
func (us UserControllers) Attention(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.AttentionReceiveStruct)); err == nil {
		results, err := users.Attention(rec, uid)
		us.Response(ctx, results, err)
	}
}

//CreateFavorites 创建收藏夹
func (us UserControllers) CreateFavorites(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.CreateFavoritesReceiveStruct)); err == nil {
		results, err := users.CreateFavorites(rec, uid)
		us.Response(ctx, results, err)
	}
}

//DeleteFavorites 获取收藏夹
func (us UserControllers) DeleteFavorites(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.DeleteFavoritesReceiveStruct)); err == nil {
		results, err := users.DeleteFavorites(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetFavoritesList 获取收藏夹列表
func (us UserControllers) GetFavoritesList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetFavoritesList(uid)
	us.Response(ctx, results, err)
}

//FavoriteVideo 收藏视频
func (us UserControllers) FavoriteVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.FavoriteVideoReceiveStruct)); err == nil {
		results, err := users.FavoriteVideo(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetFavoritesListByFavoriteVideo 获取收藏夹列表在视频页面
func (us UserControllers) GetFavoritesListByFavoriteVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetFavoritesListByFavoriteVideoReceiveStruct)); err == nil {
		results, err := users.GetFavoritesListByFavoriteVideo(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetFavoriteVideoList 获取收藏夹视频列表
func (us UserControllers) GetFavoriteVideoList(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetFavoriteVideoListReceiveStruct)); err == nil {
		results, err := users.GetFavoriteVideoList(rec)
		us.Response(ctx, results, err)
	}
}

//GetRecordList 获取历史记录
func (us UserControllers) GetRecordList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetRecordListReceiveStruct)); err == nil {
		results, err := users.GetRecordList(rec, uid)
		us.Response(ctx, results, err)
	}
}

//ClearRecord 清空历史记录
func (us UserControllers) ClearRecord(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.ClearRecord(uid)
	us.Response(ctx, results, err)
}

//DeleteRecordByID 删除历史记录根据id
func (us UserControllers) DeleteRecordByID(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.DeleteRecordByIDReceiveStruct)); err == nil {
		results, err := users.DeleteRecordByID(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetNoticeList 获取通知消息
func (us UserControllers) GetNoticeList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetNoticeListReceiveStruct)); err == nil {
		results, err := users.GetNoticeList(rec, uid)
		us.Response(ctx, results, err)
	}
}

//GetChatList 获取聊天列表
func (us UserControllers) GetChatList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	results, err := users.GetChatList(uid)
	us.Response(ctx, results, err)
}

//PersonalLetter 点击私信时触发
func (us UserControllers) PersonalLetter(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.PersonalLetterReceiveStruct)); err == nil {
		results, err := users.PersonalLetter(rec, uid)
		us.Response(ctx, results, err)
	}
}

//DeleteChatItem 删除聊天记录
func (us UserControllers) DeleteChatItem(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.DeleteChatItemReceiveStruct)); err == nil {
		results, err := users.DeleteChatItem(rec, uid)
		us.Response(ctx, results, err)
	}
}
