package users

import (
	"Go-Live/Utils/conversion"
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
	Gender    int       `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	IsVisible bool      `json:"is_visible"`
	Signature string    `json:"signature"`
}

//UserInfoResponse  生成返回用用户信息结构体
func (us *User) UserInfoResponse(token string) UserInfoResponse {
	//判断用户是否为微信用户进行图片处理
	var photo string
	if len(us.Openid) <= 0 {
		photo = conversion.FormattingSrc(us.Photo)
	} else {
		photo = us.Photo
	}
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
		Gender:    us.Gender,
		BirthDate: us.BirthDate,
		IsVisible: conversion.IntTurnBool(us.IsVisible),
		Signature: us.Signature,
	}
}
