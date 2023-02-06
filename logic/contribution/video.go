package contribution

import (
	"Go-Live/consts"
	receive "Go-Live/interaction/receive/contribution/video"
	response "Go-Live/interaction/response/contribution/video"
	"Go-Live/logic/contribution/socket"
	"Go-Live/models/common"
	"Go-Live/models/contribution/video"
	"Go-Live/models/contribution/video/barrage"
	"Go-Live/models/contribution/video/comments"
	"Go-Live/utils/conversion"
	"encoding/json"
	"fmt"
	"strconv"
)

func CreateVideoContribution(data *receive.CreateVideoContributionReceiveStruct, userID uint) (results interface{}, err error) {
	videoSrc, _ := json.Marshal(common.Img{
		Src: data.Video,
		Tp:  data.VideoUploadType,
	})
	coverImg, _ := json.Marshal(common.Img{
		Src: data.Cover,
		Tp:  data.CoverUploadType,
	})
	videoContribution := video.VideosContribution{
		Uid:           userID,
		Title:         data.Title,
		Video:         videoSrc,
		Cover:         coverImg,
		VideoDuration: data.VideoDuration,
		Reprinted:     conversion.BoolTurnInt8(*data.Reprinted),
		Timing:        conversion.BoolTurnInt8(*data.Timing),
		TimingTime:    data.TimingTime,
		Label:         conversion.MapConversionString(data.Label),
		Introduce:     data.Introduce,
		Heat:          0,
	}
	if *data.Timing {
		//发布视频后进行的推送相关（待开发）
	}
	if !videoContribution.Create() {
		return nil, fmt.Errorf("保存失败")
	}
	return "保存成功", nil
}

func GetVideoContributionByID(data *receive.GetVideoContributionByIDReceiveStruct) (results interface{}, err error) {
	videoInfo := new(video.VideosContribution)
	err = videoInfo.FindByID(data.VideoID)
	if err != nil {
		return nil, err
	}
	recommendList := new(video.VideosContributionList)
	err = recommendList.GetRecommendList()
	if err != nil {
		return nil, err
	}
	res := response.GetVideoContributionByIDResponse(videoInfo, recommendList)
	return res, nil
}

func SendVideoBarrage(data *receive.SendVideoBarrageReceiveStruct, uid uint) (results interface{}, err error) {
	//保存弹幕
	videoID, _ := strconv.ParseUint(data.ID, 0, 19)
	bg := barrage.Barrage{
		Uid:     uid,
		VideoID: uint(videoID),
		Time:    data.Time,
		Author:  data.Author,
		Type:    data.Type,
		Text:    data.Text,
		Color:   data.Color,
	}
	if !bg.Create() {
		return data, fmt.Errorf("发送弹幕失败")
	}
	//socket消息通知
	res := socket.ChanInfo{
		Type: consts.VideoSocketTypeResponseBarrageNum,
		Data: nil,
	}
	for _, v := range socket.Severe.VideoRoom[uint(videoID)] {
		v.MsgList <- res
	}

	return data, nil
}

func GetVideoBarrage(data *receive.GetVideoBarrageReceiveStruct) (results interface{}, err error) {
	//获取视频弹幕
	list := new(barrage.BarragesList)
	videoID, _ := strconv.ParseUint(data.ID, 0, 19)
	if !list.GetVideoBarrageByID(uint(videoID)) {
		return nil, fmt.Errorf("查询失败")
	}
	res := response.GetVideoBarrageResponse(list)
	return res, nil
}

func GetVideoBarrageList(data *receive.GetVideoBarrageListReceiveStruct) (results interface{}, err error) {
	//获取视频弹幕
	list := new(barrage.BarragesList)
	videoID, _ := strconv.ParseUint(data.ID, 0, 19)
	if !list.GetVideoBarrageByID(uint(videoID)) {
		return nil, fmt.Errorf("查询失败")
	}
	res := response.GetVideoBarrageListResponse(list)
	return res, nil
}

func VideoPostComment(data *receive.VideosPostCommentReceiveStruct, userID uint) (results interface{}, err error) {
	ct := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentFirstID := ct.GetCommentFirstID()

	ctu := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentUserID := ctu.GetCommentUserID()
	comment := comments.Comment{
		Uid:            userID,
		VideoID:        data.VideoID,
		Context:        data.Content,
		CommentID:      data.ContentID,
		CommentUserID:  CommentUserID,
		CommentFirstID: CommentFirstID,
	}
	if !comment.Create() {
		return nil, fmt.Errorf("发布失败")
	}
	return "发布成功", nil
}

func GetVideoComment(data *receive.GetVideoCommentReceiveStruct) (results interface{}, err error) {
	videosContribution := new(video.VideosContribution)
	if !videosContribution.GetVideoComments(data.VideoID, data.PageInfo) {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetVideoContributionCommentsResponse(videosContribution), nil
}
