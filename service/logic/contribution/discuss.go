package contribution

import (
	receive "Go-Live/interaction/receive/contribution/discuss"
	response "Go-Live/interaction/response/contribution/discuss"
	"Go-Live/models/contribution/article"
	articleComments "Go-Live/models/contribution/article/comments"
	"Go-Live/models/contribution/video"
	"Go-Live/models/contribution/video/barrage"
	videoComments "Go-Live/models/contribution/video/comments"
	"fmt"
)

func GetDiscussVideoList(data *receive.GetDiscussVideoListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取用户发布的视频
	videoList := new(video.VideosContributionList)
	err = videoList.GetDiscussVideoCommentList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询视频相关信息失败")
	}
	videoIDs := make([]uint, 0)
	for _, v := range *videoList {
		videoIDs = append(videoIDs, v.ID)
	}
	//得到视频信息
	cml := new(videoComments.CommentList)
	err = cml.GetCommentListByIDs(videoIDs, data.PageInfo)
	if err != nil {
		return nil, fmt.Errorf("查询视频评论信息失败")
	}

	return response.GetDiscussVideoListResponse(cml), nil
}

func GetDiscussArticleList(data *receive.GetDiscussArticleListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取用户发布的专栏
	articleList := new(article.ArticlesContributionList)
	err = articleList.GetDiscussArticleCommentList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询专栏相关信息失败")
	}
	articleIDs := make([]uint, 0)
	for _, v := range *articleList {
		articleIDs = append(articleIDs, v.ID)
	}
	//得到文章信息
	cml := new(articleComments.CommentList)
	err = cml.GetCommentListByIDs(articleIDs, data.PageInfo)
	if err != nil {
		return nil, fmt.Errorf("查询文章评论信息失败")
	}
	return response.GetDiscussArticleListResponse(cml), nil
}

func GetDiscussBarrageList(data *receive.GetDiscussBarrageListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取用户发布的视频
	videoList := new(video.VideosContributionList)
	err = videoList.GetDiscussVideoCommentList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询视频相关信息失败")
	}
	videoIDs := make([]uint, 0)
	for _, v := range *videoList {
		videoIDs = append(videoIDs, v.ID)
	}
	//得到视频弹幕信息
	cml := new(barrage.BarragesList)
	err = cml.GetBarrageListByIDs(videoIDs, data.PageInfo)
	if err != nil {
		return nil, fmt.Errorf("查询视频弹幕信息失败")
	}
	return response.GetDiscussBarrageListResponse(cml), nil
}
