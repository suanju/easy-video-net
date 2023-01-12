package home

import (
	"Go-Live/models/contribution/video"
	"Go-Live/models/home"
	"Go-Live/models/home/rotograph"
)

func GetHomeInfo(data *home.GetHomeInfoReceiveStruct) (results interface{}, err error) {
	//获取主页轮播图
	rotographList := new(rotograph.List)
	err = rotographList.GetAll()
	if err != nil {
		return nil, err
	}

	//获取主页推荐视频
	videoList := new(video.VideosContribution)
	err = videoList.GetHoneVideoList()
	if err != nil {
		return nil, err
	}
	response := &home.GetHomeInfoResponse{}
	response.Response(rotographList)
	return response, nil
}
