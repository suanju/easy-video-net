package users

import "time"

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

type SetUserInfoStruct struct {
	Username  string    `json:"username" binding:"required"`
	Gender    *int      `json:"gender" binding:"required"`
	BirthDate time.Time `json:"birth_Date" binding:"required"`
	IsVisible *bool     `json:"is_Visible" binding:"required"`
	Signature string    `json:"signature" binding:"required"`
}
