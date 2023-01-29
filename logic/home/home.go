package home

import (
	receive "Go-Live/interaction/receive/home"
	response "Go-Live/interaction/response/home"
	"Go-Live/models/contribution/video"
	"Go-Live/models/home/rotograph"
)

func GetHomeInfo(data *receive.GetHomeInfoReceiveStruct) (results interface{}, err error) {
	//获取主页轮播图
	rotographList := new(rotograph.List)
	err = rotographList.GetAll()
	if err != nil {
		return nil, err
	}

	//获取主页推荐视频
	videoList := new(video.VideosContributionList)
	err = videoList.GetHoneVideoList()

	if err != nil {
		return nil, err
	}
	res := &response.GetHomeInfoResponse{}
	res.Response(rotographList, videoList)

	return res, nil
}
