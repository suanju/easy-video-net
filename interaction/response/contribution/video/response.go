package response

import (
	"Go-Live/models/contribution/video"
	"Go-Live/utils/conversion"
	"time"
)

//Info 视频信息
type Info struct {
	ID            uint        `json:"id"`
	Uid           uint        `json:"uid" `
	Title         string      `json:"title" `
	Video         string      `json:"video"`
	Cover         string      `json:"cover" `
	VideoDuration int64       `json:"video_duration"`
	Label         []string    `json:"label"`
	Introduce     string      `json:"introduce"`
	Heat          int         `json:"heat"`
	BarrageNumber int         `json:"barrageNumber"`
	CreatorInfo   creatorInfo `json:"creatorInfo"`
	CreatedAt     time.Time   `json:"created_at"`
}

//创作者信息
type creatorInfo struct {
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}

//推荐视频信息
type recommendVideo struct {
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
type RecommendList []recommendVideo

type Response struct {
	VideoInfo     Info          `json:"videoInfo"`
	RecommendList RecommendList `json:"recommendList"`
}

func GetVideoContributionByIDResponse(vc *video.VideosContribution, recommendVideoList *video.VideosContributionList) Response {
	//处理视频主要信息
	creatorAvatar, _ := conversion.FormattingJsonSrc(vc.UserInfo.Photo)
	cover, _ := conversion.FormattingJsonSrc(vc.Cover)
	videoSrc, _ := conversion.FormattingJsonSrc(vc.Video)
	response := Response{
		VideoInfo: Info{
			ID:            vc.ID,
			Uid:           vc.Uid,
			Title:         vc.Title,
			Video:         videoSrc,
			Cover:         cover,
			VideoDuration: vc.VideoDuration,
			Label:         conversion.StringConversionMap(vc.Label),
			Introduce:     vc.Introduce,
			Heat:          vc.Heat,
			BarrageNumber: len(vc.Barrage),
			CreatorInfo: creatorInfo{
				Username:  vc.UserInfo.Username,
				Avatar:    creatorAvatar,
				Signature: vc.UserInfo.Signature,
			},
			CreatedAt: vc.CreatedAt,
		},
	}
	//处理推荐视频
	rl := make(RecommendList, 0)
	for _, lk := range *recommendVideoList {
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		info := recommendVideo{
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
		rl = append(rl, info)
	}
	response.RecommendList = rl
	return response
}
