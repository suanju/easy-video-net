package users

import (
	receive "Go-Live/interaction/receive/users"
	response "Go-Live/interaction/response/users"
	"Go-Live/models/contribution/article"
	"Go-Live/models/contribution/video"
	"Go-Live/models/users"
	"Go-Live/models/users/attention"
	"fmt"
)

func GetSpaceIndividual(data *receive.GetSpaceIndividualReceiveStruct, userID uint) (results interface{}, err error) {
	//获取游戏信息
	user := new(users.User)
	user.Find(data.ID)
	isAttention := false
	at := new(attention.Attention)
	if userID != 0 {
		//获取是否关注
		isAttention = at.IsAttention(userID, data.ID)
	}
	//获取关注和粉丝
	attentionNum, err := at.GetAttentionNum(data.ID)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	vermicelliNum, err := at.GetVermicelliNum(data.ID)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetSpaceIndividualResponse(user, isAttention, attentionNum, vermicelliNum)
}

func GetReleaseInformation(data *receive.GetReleaseInformationReceiveStruct) (results interface{}, err error) {
	//获取视频列表
	videoList := new(video.VideosContributionList)
	err = videoList.GetVideoListBySpace(data.ID)
	if err != nil {
		return nil, fmt.Errorf("查询空间视频失败")
	}
	articleList := new(article.ArticlesContributionList)
	err = articleList.GetArticleBySpace(data.ID)
	if err != nil {
		return nil, fmt.Errorf("查询空间专栏失败")
	}
	res, err := response.GetReleaseInformationResponse(videoList, articleList)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}
