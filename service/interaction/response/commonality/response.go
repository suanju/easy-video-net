package commonality

import (
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/commonality"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/users"
	"easy-video-net/utils/calculate"
	"easy-video-net/utils/conversion"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"time"
)

func UploadingMethodResponse(tp string) interface{} {
	type UploadingMethodResponseStruct struct {
		Tp string `json:"type"`
	}
	return UploadingMethodResponseStruct{
		Tp: tp,
	}
}

func UploadingDirResponse(dir string) interface{} {
	type UploadingDirResponseStruct struct {
		Path string `json:"path"`
	}
	return UploadingDirResponseStruct{
		Path: dir,
	}
}

//VideoInfo 首页视频
type VideoInfo struct {
	ID            uint      `json:"id"`
	Uid           uint      `json:"uid" `
	Title         string    `json:"title" `
	Video         string    `json:"video"`
	Cover         string    `json:"cover" `
	VideoDuration int64     `json:"video_duration"`
	Label         []string  `json:"label"`
	Introduce     string    `json:"introduce"`
	Heat          int       `json:"heat"`
	BarrageNumber int       `json:"barrageNumber"`
	Username      string    `json:"username"`
	CreatedAt     time.Time `json:"created_at"`
}

type videoInfoList []VideoInfo

func SearchVideoResponse(videoList *video.VideosContributionList) (interface{}, error) {
	//处理视频
	vl := make(videoInfoList, 0)
	for _, lk := range *videoList {
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		info := VideoInfo{
			ID:            lk.ID,
			Uid:           lk.Uid,
			Title:         lk.Title,
			Video:         videoSrc,
			Cover:         cover,
			VideoDuration: lk.VideoDuration,
			Label:         conversion.StringConversionMap(lk.Label),
			Introduce:     lk.Introduce,
			Heat:          lk.Heat,
			BarrageNumber: len(lk.Barrage),
			Username:      lk.UserInfo.Username,
			CreatedAt:     lk.CreatedAt,
		}
		vl = append(vl, info)
	}
	return vl, nil
}

type UserInfo struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Photo       string `json:"photo"`
	Signature   string `json:"signature"`
	IsAttention bool   `json:"is_attention"`
}

type UserInfoList []UserInfo

func SearchUserResponse(userList *users.UserList, aids []uint) (interface{}, error) {
	list := make(UserInfoList, 0)
	for _, v := range *userList {
		photo, _ := conversion.FormattingJsonSrc(v.Photo)
		list = append(list, UserInfo{
			ID:          v.ID,
			Username:    v.Username,
			Photo:       photo,
			Signature:   v.Signature,
			IsAttention: calculate.ArrayIsContain(aids, v.ID),
		})
	}
	return list, nil
}

type UploadCheckStruct struct {
	IsUpload bool                    `json:"is_upload"`
	List     receive.UploadSliceList `json:"list"`
	Path     string                  `json:"path"`
}

func UploadCheckResponse(is bool, list receive.UploadSliceList, path string) (interface{}, error) {
	return UploadCheckStruct{
		IsUpload: is,
		List:     list,
		Path:     path,
	}, nil
}

type GteStsInfoStruct struct {
	Region          string `json:"region"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	StsToken        string `json:"sts_token"`
	Bucket          string `json:"bucket"`
	ExpirationTime  int64  `json:"expiration_time"`
}

func GteStsInfo(info *sts20150401.AssumeRoleResponseBodyCredentials) (interface{}, error) {
	return GteStsInfoStruct{
		Region:          global.Config.AliyunOss.Region,
		AccessKeyID:     *info.AccessKeyId,
		AccessKeySecret: *info.AccessKeySecret,
		StsToken:        *info.SecurityToken,
		Bucket:          global.Config.AliyunOss.Bucket,
		ExpirationTime:  time.Now().Unix() + int64(global.Config.AliyunOss.DurationSeconds),
	}, nil
}
