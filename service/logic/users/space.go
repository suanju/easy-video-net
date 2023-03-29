package users

import (
	receive "easy-video-net/interaction/receive/users"
	response "easy-video-net/interaction/response/users"
	"easy-video-net/models/contribution/article"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/users"
	"easy-video-net/models/users/attention"
	"fmt"
)

func GetSpaceIndividual(data *receive.GetSpaceIndividualReceiveStruct, uid uint) (results interface{}, err error) {
	//获取游戏信息
	user := new(users.User)
	user.Find(data.ID)
	isAttention := false
	at := new(attention.Attention)
	if uid != 0 {
		//获取是否关注
		isAttention = at.IsAttention(uid, data.ID)
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

func GetAttentionList(data *receive.GetAttentionListReceiveStruct, id uint) (results interface{}, err error) {
	//获取用户关注列表
	al := new(attention.AttentionsList)
	err = al.GetAttentionList(data.ID)
	if err != nil {
		return nil, fmt.Errorf("获取失败")
	}
	//获取自己关注的用户
	ual := new(attention.AttentionsList)
	arr, err := ual.GetAttentionListByIdArr(id)
	if err != nil {
		return nil, fmt.Errorf("获取失败")
	}
	res, err := response.GetAttentionListResponse(al, arr)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}

func GetVermicelliList(data *receive.GetVermicelliListReceiveStruct, id uint) (results interface{}, err error) {
	//获取用户粉丝列表
	al := new(attention.AttentionsList)
	err = al.GetVermicelliList(data.ID)
	if err != nil {
		return nil, fmt.Errorf("获取失败")
	}
	//获取自己关注的用户
	ual := new(attention.AttentionsList)
	arr, err := ual.GetAttentionListByIdArr(id)
	if err != nil {
		return nil, fmt.Errorf("获取失败")
	}
	res, err := response.GetVermicelliListResponse(al, arr)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}
