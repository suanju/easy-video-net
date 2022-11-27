package users

import (
	"Go-Live/utils/conversion"
	"time"
)

type UserInfoResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Photo    string `json:"photo"`
	Token    string `json:"token"`
}

type UserSetInfoResponse struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Gender    int8      `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	IsVisible bool      `json:"is_visible"`
	Signature string    `json:"signature"`
}

//UserInfoResponse  生成返回用用户信息结构体
func (us *User) UserInfoResponse(token string) UserInfoResponse {
	//判断用户是否为微信用户进行图片处理
	var photo string
	photo, _ = conversion.FormattingJsonSrc(us.Photo)
	return UserInfoResponse{
		ID:       us.ID,
		UserName: us.Username,
		Photo:    photo,
		Token:    token,
	}
}

//UserSetInfoResponse  设置用户信息返回结果
func (us *User) UserSetInfoResponse() UserSetInfoResponse {
	return UserSetInfoResponse{
		ID:        us.ID,
		UserName:  us.Username,
		Gender:    int8(us.Gender),
		BirthDate: us.BirthDate,
		IsVisible: conversion.IntTurnBool(int(us.IsVisible)),
		Signature: us.Signature,
	}
}
