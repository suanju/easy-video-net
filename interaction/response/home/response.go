package response

import (
	"Go-Live/models/contribution/video"
	"Go-Live/models/home/rotograph"
	"Go-Live/utils/conversion"
	"time"
)

//首页轮播图
type rotographInfo struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Color string `json:"color"`
	Type  string `json:"type"`
	ToId  uint   `json:"to_id"`
}
type rotographInfoList []rotographInfo

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

type GetHomeInfoResponse struct {
	Rotograph rotographInfoList `json:"rotograph"`
	VideoList videoInfoList     `json:"videoList"`
}

func (r *GetHomeInfoResponse) Response(rotographList *rotograph.List, videoList *video.VideosContributionList) {
	//处理轮播图
	rl := make(rotographInfoList, 0)
	for _, rk := range *rotographList {
		cover, _ := conversion.FormattingJsonSrc(rk.Cover)
		info := rotographInfo{
			Title: rk.Title,
			Cover: cover,
			Color: rk.Color,
			Type:  rk.Type,
			ToId:  rk.ToId,
		}
		rl = append(rl, info)
	}
	r.Rotograph = rl
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
	r.VideoList = vl

}
