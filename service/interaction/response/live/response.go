package response

import (
	"Go-Live/models/users"
	"Go-Live/utils/conversion"
)

type GetLiveRoomResponseStruct struct {
	Address string `json:"address"`
	Key     string `json:"key"`
}

func GetLiveRoomResponse(address string, key string) interface{} {
	return GetLiveRoomResponseStruct{
		Address: address,
		Key:     key,
	}
}

type GetLiveRoomInfoResponseStruct struct {
	Username  string `json:"username"`
	Photo     string `json:"photo"`
	LiveTitle string `json:"live_title"`
	Flv       string `json:"flv"`
}

func GetLiveRoomInfoResponse(info *users.User, flv string) interface{} {
	photo, _ := conversion.FormattingJsonSrc(info.Photo)
	return GetLiveRoomInfoResponseStruct{
		Username:  info.Username,
		Photo:     photo,
		LiveTitle: info.LiveInfo.Title,
		Flv:       flv,
	}
}

type BeLiveInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Online   int    `json:"online"`
}

type BeLiveInfoList []BeLiveInfo

func GetBeLiveListResponse(ul *users.UserList) interface{} {
	list := make(BeLiveInfoList, 0)
	for _, v := range *ul {
		photo, _ := conversion.FormattingJsonSrc(v.Photo)
		img, _ := conversion.FormattingJsonSrc(v.LiveInfo.Img)
		list = append(list, BeLiveInfo{
			ID:       v.ID,
			Username: v.Username,
			Photo:    photo,
			Img:      img,
			Title:    v.LiveInfo.Title,
			Online:   0,
		})
	}
	return list
}
