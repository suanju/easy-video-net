package contribution

import (
	receive "Go-Live/interaction/receive/contribution/discuss"
	"Go-Live/models/contribution/video"
	"fmt"
)

func GetDiscussVideoList(data *receive.GetDiscussVideoListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取用户发布的视频
	videoList := new(video.VideosContributionList)
	err = videoList.GetDiscussVideoCommentList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询视频评论信息失败")
	}
	videoIDs := make([]uint, 0)
	for _, v := range *videoList {
		videoIDs = append(videoIDs, v.ID)
	}
	return videoList, nil
}
