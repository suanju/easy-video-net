package response

import (
	"Go-Live/models/users"
	"Go-Live/utils/conversion"
)

type GetLiveRoomResponseStruct struct {
	Address string `json:"address"`
	Key     string `json:"key"`
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
