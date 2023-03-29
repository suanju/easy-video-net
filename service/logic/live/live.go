package live

import (
	"Go-Live/global"
	receive "Go-Live/interaction/receive/live"
	response "Go-Live/interaction/response/live"
	"Go-Live/models/users"
	"Go-Live/models/users/record"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetLiveRoom(uid uint) (results interface{}, err error) {
	//请求直播服务器
	url := global.Config.LiveConfig.Agreement + "://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.Api + "/control/get?room="
	url = url + strconv.Itoa(int(uid))
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
	return response.GetLiveRoomResponse("rtmp://"+global.Config.LiveConfig.IP+":"+global.Config.LiveConfig.RTMP+"/live", ReqGetRoom.Data), nil
}

func GetLiveRoomInfo(data *receive.GetLiveRoomInfoReceiveStruct, uid uint) (results interface{}, err error) {
	userInfo := new(users.User)
	userInfo.FindLiveInfo(data.RoomID)
	flv := global.Config.LiveConfig.Agreement + "://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.FLV + "/live/" + strconv.Itoa(int(data.RoomID)) + ".flv"

	if uid > 0 {
		//添加历史记录
		rd := new(record.Record)
		err = rd.AddLiveRecord(uid, data.RoomID)
		if err != nil {
			return nil, fmt.Errorf("添加历史记录失败")
		}
	}
	return response.GetLiveRoomInfoResponse(userInfo, flv), nil
}

func GetBeLiveList() (results interface{}, err error) {
	//取开通播放用户id
	url := global.Config.LiveConfig.Agreement + "://" + global.Config.LiveConfig.IP + ":" + global.Config.LiveConfig.Api + "/stat/livestat"
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
	livestat := new(receive.LivestatRes)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(livestat); err != nil {
		return nil, fmt.Errorf("解析失败")
	}
	if livestat.Status != 200 {
		return nil, fmt.Errorf("获取直播列表失败")
	}
	//获取live中正在值得列表
	keys := make([]uint, 0)
	for _, kv := range livestat.Data.Publishers {
		ka := strings.Split(kv.Key, "live/")
		uintKey, _ := strconv.ParseUint(ka[1], 10, 19)
		keys = append(keys, uint(uintKey))
	}
	userList := new(users.UserList)
	if len(keys) > 0 {
		err = userList.GetBeLiveList(keys)
		if err != nil {
			return nil, fmt.Errorf("查询失败")
		}
	}
	return response.GetBeLiveListResponse(userList), nil
}
