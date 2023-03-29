package receive

import (
	"easy-video-net/models/common"
	"time"
)

// WxAuthorizationReceiveStruct 微信登入接口数据结构体
type WxAuthorizationReceiveStruct struct {
	AvatarUrl string `json:"avatarUrl" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Gender    string `json:"gender" `
	NickName  string `json:"nickName" binding:"required"`
}

// RegisterReceiveStruct 用户注册
type RegisterReceiveStruct struct {
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

//LoginReceiveStruct 用户登入
type LoginReceiveStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//SendEmailVerCodeReceiveStruct 邮箱验证码
type SendEmailVerCodeReceiveStruct struct {
	Email string `json:"email" binding:"required,email"`
}

// ForgetReceiveStruct 用户找回密码
type ForgetReceiveStruct struct {
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}
type DetermineNameExistsStruct struct {
	Username string `json:"username" binding:"required"`
}

//SetUserInfoReceiveStruct 设置用户信息
type SetUserInfoReceiveStruct struct {
	Username  string    `json:"username" binding:"required"`
	Gender    *int      `json:"gender" binding:"required"`
	BirthDate time.Time `json:"birth_Date" binding:"required"`
	IsVisible *bool     `json:"is_Visible" binding:"required"`
	Signature string    `json:"signature" binding:"required"`
}

//UpdateAvatarStruct 更新头像
type UpdateAvatarStruct struct {
	ImgUrl string `json:"imgUrl" binding:"required"`
	Tp     string `json:"type" binding:"required"`
}

//SaveLiveDataReceiveStruct 设置直播信息
type SaveLiveDataReceiveStruct struct {
	Tp     string `json:"type" binding:"required"`
	ImgUrl string `json:"imgUrl" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type ChangePasswordReceiveStruct struct {
	VerificationCode string `json:"verificationCode" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ConfirmPassword  string `json:"confirm_password" binding:"required"`
}

type AttentionReceiveStruct struct {
	Uid uint `json:"uid"  binding:"required" binding:"required"`
}

type GetSpaceIndividualReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type GetReleaseInformationReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type GetAttentionListReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type GetVermicelliListReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteFavoritesReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type CreateFavoritesReceiveStruct struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	Tp      string `json:"type"`
}

type FavoriteVideoReceiveStruct struct {
	IDs     []uint `json:"ids" binding:"required"`
	VideoID uint   `json:"video_id" binding:"required"`
}

type GetFavoritesListByFavoriteVideoReceiveStruct struct {
	VideoID uint `json:"video_id" binding:"required"`
}

type GetFavoriteVideoListReceiveStruct struct {
	FavoriteID uint `json:"favorite_id" binding:"required"`
}

type GetRecordListReceiveStruct struct {
	PageInfo common.PageInfo `json:"page_info" binding:"required"`
}

type DeleteRecordByIDReceiveStruct struct {
	ID uint `json:"id"`
}

type GetNoticeListReceiveStruct struct {
	Type     string          `json:"type"`
	PageInfo common.PageInfo `json:"page_info" binding:"required"`
}

type GetChatHistoryMsgStruct struct {
	Tid      uint      `json:"tid"`
	LastTime time.Time `json:"last_time"`
}

type PersonalLetterReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteChatItemReceiveStruct struct {
	ID uint `json:"id" binding:"required"`
}
