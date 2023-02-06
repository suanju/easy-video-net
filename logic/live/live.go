package live

import (
	"Go-Live/global"
	receive "Go-Live/interaction/receive/live"
	response "Go-Live/interaction/response/live"
	"Go-Live/models/users"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetLiveRoom(data *receive.GetLiveRoomReceiveStruct, userId uint) (results interface{}, err error) {
	//请求直播服务器
	url := "http://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.GetRoom + "/control/get?room="
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = url + strconv.Itoa(int(userId))
	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	// 解析http请求中body 数据到我们定义的结构体中
	ReqGetRoom := new(receive.ReqGetRoom)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(ReqGetRoom); err != nil {
		return nil, err
	}
	if ReqGetRoom.Status != 200 {
		return nil, fmt.Errorf("获取直播地址失败")
	}
	res := response.GetLiveRoomResponseStruct{
		Address: "rtmp://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.RTMP + "/live",
		Key:     ReqGetRoom.Data,
	}
	return res, nil
}

func GetLiveRoomInfo(data *receive.GetLiveRoomInfoReceiveStruct) (results interface{}, err error) {
	userInfo := new(users.User)
	userInfo.FindLiveInfo(data.RoomID)
	flv := global.Config.LiveConfig.Agreement + "://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.FLV + "/live/" + strconv.Itoa(int(data.RoomID)) + ".flv"
	return response.GetLiveRoomInfoResponse(userInfo, flv), nil
}
