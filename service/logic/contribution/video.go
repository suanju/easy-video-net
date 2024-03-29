package contribution

import (
	"easy-video-net/consts"
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/contribution/video"
	response "easy-video-net/interaction/response/contribution/video"
	"easy-video-net/logic/contribution/sokcet"
	"easy-video-net/logic/users/notice"
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/contribution/video/barrage"
	"easy-video-net/models/contribution/video/comments"
	"easy-video-net/models/contribution/video/like"
	transcodingTask "easy-video-net/models/sundry/transcoding"
	"easy-video-net/models/users/attention"
	"easy-video-net/models/users/collect"
	"easy-video-net/models/users/favorites"
	noticeModel "easy-video-net/models/users/notice"
	"easy-video-net/models/users/record"
	"easy-video-net/utils/calculate"
	"easy-video-net/utils/conversion"
	"easy-video-net/utils/oss"
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func CreateVideoContribution(data *receive.CreateVideoContributionReceiveStruct, uid uint) (results interface{}, err error) {
	//发布视频
	videoSrc, _ := json.Marshal(common.Img{
		Src: data.Video,
		Tp:  data.VideoUploadType,
	})
	coverImg, _ := json.Marshal(common.Img{
		Src: data.Cover,
		Tp:  data.CoverUploadType,
	})
	var width, height int
	if data.VideoUploadType == "local" {
		//如果是本地上传
		width, height, err = calculate.GetVideoResolution(data.Video)
		if err != nil {
			global.Logger.Error("获取视频分辨率失败")
			return nil, fmt.Errorf("获取视频分辨率失败")
		}
	} else {
		mediaInfo, err := oss.GetMediaInfo(data.Media)
		if err != nil {
			return nil, fmt.Errorf("获取视频信息失败,稍后再试")
		}
		width, _ = strconv.Atoi(*mediaInfo.Body.MediaInfo.FileInfoList[0].FileBasicInfo.Width)
		height, _ = strconv.Atoi(*mediaInfo.Body.MediaInfo.FileInfoList[0].FileBasicInfo.Height)
	}
	videoContribution := &video.VideosContribution{
		Uid:           uid,
		Title:         data.Title,
		Cover:         coverImg,
		Reprinted:     conversion.BoolTurnInt8(*data.Reprinted),
		Label:         conversion.MapConversionString(data.Label),
		VideoDuration: data.VideoDuration,
		Introduce:     data.Introduce,
		Heat:          0,
	}
	if data.Media != nil {
		videoContribution.MediaID = *data.Media
	}
	// 定义转码分辨率列表
	resolutions := []int{1080, 720, 480, 360}
	if height >= 1080 {
		resolutions = resolutions[1:]
		videoContribution.Video = videoSrc
	} else if height >= 720 && height < 1080 {
		resolutions = resolutions[2:]
		videoContribution.Video720p = videoSrc
	} else if height >= 480 && height < 720 {
		resolutions = resolutions[3:]
		videoContribution.Video480p = videoSrc
	} else if height >= 360 && height < 480 {
		resolutions = resolutions[4:]
		videoContribution.Video360p = videoSrc
	} else {
		global.Logger.Error("上传视频分辨率过低")
		return nil, fmt.Errorf("上传视频分辨率过低")
	}
	if !videoContribution.Create() {
		return nil, fmt.Errorf("保存失败")
	}
	//进行视频转码
	go func(width, height int, video *video.VideosContribution) {
		if data.VideoUploadType == "local" {
			//本地ffmpeg 处理
			inputFile := data.Video
			sr := strings.Split(inputFile, ".")
			for _, r := range resolutions {
				// 计算转码后的宽和高需要取整
				w := int(math.Ceil(float64(r) / float64(height) * float64(width)))
				h := int(math.Ceil(float64(r)))
				if h >= height {
					continue
				}
				dst := sr[0] + fmt.Sprintf("_output_%dp."+sr[1], r)
				// TODO: 调用转码接口，将转码后的视频保存到指定目录
				cmd := exec.Command("ffmpeg",
					"-i",
					inputFile,
					"-vf",
					fmt.Sprintf("scale=%d:%d", w, h),
					"-c:a",
					"copy",
					"-c:v",
					"libx264",
					"-preset",
					"medium",
					"-crf",
					"23",
					"-y",
					dst)
				err = cmd.Run()
				if err != nil {
					global.Logger.Errorf("视频 :%s :转码%d*%d失败 cmd : %s ,err :%s", inputFile, w, h, cmd, err)
					continue
				}
				src, _ := json.Marshal(common.Img{
					Src: dst,
					Tp:  "local",
				})
				switch r {
				case 1080:
					videoContribution.Video = src
				case 720:
					videoContribution.Video720p = src
				case 480:
					videoContribution.Video480p = src
				case 360:
					videoContribution.Video360p = src
				}
				if !videoContribution.Save() {
					global.Logger.Errorf("视频 :%s : 转码%d*%d后视频保存到数据库失败", inputFile, w, h)
				}
				global.Logger.Infof("视频 :%s : 转码%d*%d成功", inputFile, w, h)
			}
		} else if data.VideoUploadType == "aliyunOss" && global.Config.AliyunOss.IsOpenTranscoding {
			inputFile := data.Video
			sr := strings.Split(inputFile, ".")
			//云转码处理
			for _, r := range resolutions {
				//获取转码模板
				var template string
				dst := sr[0] + fmt.Sprintf("_output_%dp."+sr[1], r)
				src, _ := json.Marshal(common.Img{
					Src: dst,
					Tp:  data.VideoUploadType,
				})
				switch r {
				case 1080:
					template = global.Config.AliyunOss.TranscodingTemplate1080p
					videoContribution.Video = src
				case 720:
					template = global.Config.AliyunOss.TranscodingTemplate720p
					videoContribution.Video720p = src
				case 480:
					template = global.Config.AliyunOss.TranscodingTemplate480p
					videoContribution.Video480p = src
				case 360:
					template = global.Config.AliyunOss.TranscodingTemplate360p
					videoContribution.Video360p = src
				}
				outputUrl, _ := conversion.SwitchIngStorageFun(data.VideoUploadType, dst)
				taskName := "转码 : " + *data.Media + "时间 :" + time.Now().Format("2006.01.02 15:04:05") + " template : " + template
				jobInfo, err := oss.SubmitTranscodeJob(taskName, video.MediaID, outputUrl, template)
				if err != nil {
					global.Logger.Errorf("视频云转码 : %s 失败 err : %s", outputUrl, err.Error())
					continue
				}
				task := &transcodingTask.TranscodingTask{
					TaskID:     *jobInfo.TranscodeParentJob.ParentJobId,
					VideoID:    video.ID,
					Resolution: r,
					Dst:        dst,
					Status:     0,
					Type:       transcodingTask.Aliyun,
				}
				if !task.AddTask() {
					global.Logger.Errorf("视频云转码任务名: %s 后将视频任务 保存到数据库失败", taskName)
				}
			}
		}
	}(width, height, videoContribution)

	return "保存成功", nil
}

func UpdateVideoContribution(data *receive.UpdateVideoContributionReceiveStruct, uid uint) (results interface{}, err error) {
	//更新视频
	videoInfo := new(video.VideosContribution)
	err = videoInfo.FindByID(data.ID)
	if err != nil {
		return nil, fmt.Errorf("操作视频不存在")
	}
	if videoInfo.Uid != uid {
		return nil, fmt.Errorf("非法操作")
	}
	coverImg, _ := json.Marshal(common.Img{
		Src: data.Cover,
		Tp:  data.CoverUploadType,
	})
	updateList := map[string]interface{}{
		"cover":     coverImg,
		"title":     data.Title,
		"label":     conversion.MapConversionString(data.Label),
		"reprinted": conversion.BoolTurnInt8(*data.Reprinted),
		"introduce": data.Introduce,
	}
	//进行视频资料更新
	if !videoInfo.Update(updateList) {
		return nil, fmt.Errorf("更新数据失败")
	}
	return "更新成功", nil
}

func DeleteVideoByID(data *receive.DeleteVideoByIDReceiveStruct, uid uint) (results interface{}, err error) {
	vo := new(video.VideosContribution)
	if !vo.Delete(data.ID, uid) {
		return nil, fmt.Errorf("删除失败")
	}
	return "删除成功", nil
}

func GetVideoContributionByID(data *receive.GetVideoContributionByIDReceiveStruct, uid uint) (results interface{}, err error) {
	videoInfo := new(video.VideosContribution)
	//获取视频信息
	err = videoInfo.FindByID(data.VideoID)
	if err != nil {
		return nil, fmt.Errorf("查询信息失败")
	}
	isAttention := false
	isLike := false
	isCollect := false
	if uid != 0 {
		//进行视频播放增加
		if !global.RedisDb.SIsMember(consts.VideoWatchByID+strconv.Itoa(int(data.VideoID)), uid).Val() {
			//最近无播放
			global.RedisDb.SAdd(consts.VideoWatchByID+strconv.Itoa(int(data.VideoID)), uid)
			if videoInfo.Watch(data.VideoID) != nil {
				global.Logger.Error("添加播放量错误视频video_id:", videoInfo.Watch(data.VideoID))
			}
			videoInfo.Heat++
		}
		//获取是否关注
		at := new(attention.Attention)
		isAttention = at.IsAttention(uid, videoInfo.UserInfo.ID)

		//获取是否关注
		lk := new(like.Likes)
		isLike = lk.IsLike(uid, videoInfo.ID)

		//判断是否已经收藏
		fl := new(favorites.FavoriteList)
		err = fl.GetFavoritesList(uid)
		if err != nil {
			return nil, fmt.Errorf("查询失败")
		}
		flIDs := make([]uint, 0)
		for _, v := range *fl {
			flIDs = append(flIDs, v.ID)
		}
		//判断是否在收藏夹内
		cl := new(collect.CollectsList)
		isCollect = cl.FindIsCollectByFavorites(data.VideoID, flIDs)

		//添加历史记录
		rd := new(record.Record)
		err = rd.AddVideoRecord(uid, data.VideoID)
		if err != nil {
			return nil, fmt.Errorf("添加历史记录失败")
		}

	}
	//获取推荐列表
	recommendList := new(video.VideosContributionList)
	err = recommendList.GetRecommendList()
	if err != nil {
		return nil, err
	}
	res := response.GetVideoContributionByIDResponse(videoInfo, recommendList, isAttention, isLike, isCollect)
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
	res := sokcet.ChanInfo{
		Type: consts.VideoSocketTypeResponseBarrageNum,
		Data: nil,
	}
	for _, v := range sokcet.Severe.VideoRoom[uint(videoID)] {
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

func VideoPostComment(data *receive.VideosPostCommentReceiveStruct, uid uint) (results interface{}, err error) {
	videoInfo := new(video.VideosContribution)
	err = videoInfo.FindByID(data.VideoID)
	if err != nil {
		return nil, fmt.Errorf("视频不存在")
	}

	ct := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentFirstID := ct.GetCommentFirstID()

	ctu := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentUserID := ctu.GetCommentUserID()
	comment := comments.Comment{
		Uid:            uid,
		VideoID:        data.VideoID,
		Context:        data.Content,
		CommentID:      data.ContentID,
		CommentUserID:  CommentUserID,
		CommentFirstID: CommentFirstID,
	}
	if !comment.Create() {
		return nil, fmt.Errorf("发布失败")
	}

	//socket推送(在线的情况下)
	if _, ok := notice.Severe.UserMapChannel[videoInfo.UserInfo.ID]; ok {
		userChannel := notice.Severe.UserMapChannel[videoInfo.UserInfo.ID]
		userChannel.NoticeMessage(noticeModel.VideoComment)
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

func GetVideoManagementList(data *receive.GetVideoManagementListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取个人发布视频信息
	list := new(video.VideosContributionList)
	err = list.GetVideoManagementList(data.PageInfo, uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	res, err := response.GetVideoManagementListResponse(list)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}

func LikeVideo(data *receive.LikeVideoReceiveStruct, uid uint) (results interface{}, err error) {
	//点赞视频
	videoInfo := new(video.VideosContribution)
	err = videoInfo.FindByID(data.VideoID)
	if err != nil {
		return nil, fmt.Errorf("视频不存在")
	}
	lk := new(like.Likes)
	err = lk.Like(uid, data.VideoID, videoInfo.UserInfo.ID)
	if err != nil {
		return nil, fmt.Errorf("操作失败")
	}

	//socket推送(在线的情况下)
	if _, ok := notice.Severe.UserMapChannel[videoInfo.UserInfo.ID]; ok {
		userChannel := notice.Severe.UserMapChannel[videoInfo.UserInfo.ID]
		userChannel.NoticeMessage(noticeModel.VideoLike)
	}

	return "操作成功", nil
}
