package home

import (
	"Go-Live/models/home/rotograph"
	"Go-Live/utils/conversion"
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

//首页视频

type VideoInfo struct {
	Uid           uint   `json:"uid" `
	Title         string `json:"title" `
	Video         string `json:"video"`
	Cover         string `json:"cover" gorm:"cover"`
	VideoDuration int64  `json:"video_duration" gorm:"video_duration"`
	Label         string `json:"label" gorm:"label"`
	Introduce     string `json:"introduce" gorm:"introduce"`
	Heat          int    `json:"heat" gorm:"heat"`
}

type videoInfoList []VideoInfo

type GetHomeInfoResponse struct {
	Rotograph rotographInfoList `json:"rotograph"`
	VideoList videoInfoList     `json:"videoList"`
}

func (r *GetHomeInfoResponse) Response(rotographList *rotograph.List) {
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
}
