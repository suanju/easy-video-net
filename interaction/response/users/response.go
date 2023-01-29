package response

import (
	"Go-Live/models/users"
	"Go-Live/models/users/liveInfo"
	"Go-Live/utils/conversion"
	"fmt"
	"time"
)

type UserInfoResponseStruct struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Photo    string `json:"photo"`
	Token    string `json:"token"`
}

//UserInfoResponse  生成返回用用户信息结构体
func UserInfoResponse(us *users.User, token string) UserInfoResponseStruct {
	//判断用户是否为微信用户进行图片处理
	var photo string
	photo, _ = conversion.FormattingJsonSrc(us.Photo)
	return UserInfoResponseStruct{
		ID:       us.ID,
		UserName: us.Username,
		Photo:    photo,
		Token:    token,
	}
}

type UserSetInfoResponseStruct struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Gender    int8      `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	IsVisible bool      `json:"is_visible"`
	Signature string    `json:"signature"`
}

//UserSetInfoResponse  设置用户信息返回结果
func UserSetInfoResponse(us *users.User) UserSetInfoResponseStruct {
	return UserSetInfoResponseStruct{
		ID:        us.ID,
		UserName:  us.Username,
		Gender:    int8(us.Gender),
		BirthDate: us.BirthDate,
		IsVisible: conversion.IntTurnBool(int(us.IsVisible)),
		Signature: us.Signature,
	}
}

type GetLiveDataResponseStruct struct {
	Img   string `json:"img"`
	Title string `json:"title"`
}

//GetLiveDataResponse 响应设置信息
func GetLiveDataResponse(li *liveInfo.LiveInfo) (data any, err error) {
	src, errs := conversion.FormattingJsonSrc(li.Img)
	if errs != nil {
		return nil, fmt.Errorf("json format error")
	}
	return GetLiveDataResponseStruct{
		Img:   src,
		Title: li.Title,
	}, nil
}
