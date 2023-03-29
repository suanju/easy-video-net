package home

import (
	receive "easy-video-net/interaction/receive/home"
	response "easy-video-net/interaction/response/home"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/home/rotograph"
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
	err = videoList.GetHoneVideoList(data.PageInfo)

	if err != nil {
		return nil, err
	}
	res := &response.GetHomeInfoResponse{}
	res.Response(rotographList, videoList)

	return res, nil
}
