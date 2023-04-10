package callback

import (
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/callback"
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/video"
	transcodingTask "easy-video-net/models/sundry/transcoding"
	"encoding/json"
)

func AliyunTranscodingMedia(data *receive.AliyunMediaCallback[receive.AliyunTranscodingMediaStruct]) (results interface{}, err error) {
	//查询任务
	taskInfo := new(transcodingTask.TranscodingTask)
	err = taskInfo.GetInfoByTaskID(data.MessageBody.ParentJobId)
	if err != nil {
		global.Logger.Errorf("阿里云媒体服务视频转码回调失败，任务id %s", data.MessageBody.Jobs)
		return nil, nil

	}
	if taskInfo.ID <= 0 {
		global.Logger.Errorf("阿里云媒体服务视频转码回调失败 : 任务不存在，任务id %s", data.MessageBody.Jobs)
		return nil, nil
	}
	if taskInfo.Status == 1 {

		global.Logger.Infof("阿里云媒体服务视频转码回调已经处理无需处理，任务id %s", data.MessageBody.Jobs)
		return nil, nil
	}
	if data.MessageBody.Status != "Success" {
		taskInfo.Status = 2
		if !taskInfo.Save() {
			global.Logger.Errorf("阿里云媒体服务回调成功后更新任务信息失败，任务id %s", data.MessageBody.Jobs)
		}
		global.Logger.Errorf("阿里云媒体服务视频转码失败，任务id %s", data.MessageBody.Jobs)
		return nil, nil
	}
	videoInfo := new(video.VideosContribution)
	err = videoInfo.FindByID(taskInfo.VideoID)
	if err != nil {
		global.Logger.Errorf("阿里云媒体服务回调获取视频信息失败，任务id %s", data.MessageBody.Jobs)
		return nil, nil
	}
	if videoInfo.ID <= 0 {
		global.Logger.Errorf("阿里云媒体服务回调视频已经删除，任务id %s", data.MessageBody.Jobs)
		return nil, nil
	}
	src, _ := json.Marshal(common.Img{
		Src: taskInfo.Dst,
		Tp:  "aliyunOss",
	})
	switch taskInfo.Resolution {
	case 1080:
		videoInfo.Video = src
	case 720:
		videoInfo.Video720p = src
	case 480:
		videoInfo.Video480p = src
	case 360:
		videoInfo.Video360p = src
	}
	if !videoInfo.Save() {
		global.Logger.Errorf("阿里云媒体服务回调成功 保存视频信息失败，任务id %s", data.MessageBody.Jobs)
	}
	taskInfo.Status = 1
	if !taskInfo.Save() {
		global.Logger.Errorf("阿里云媒体服务回调成功后更新任务信息失败，任务id %s", data.MessageBody.Jobs)
	}

	return nil, nil
}
