package home

import (
	"Go-Live/models/home/rotograph"
	"Go-Live/utils/conversion"
)

type rotographInfo struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Color string `json:"color"`
	Type  string `json:"type"`
	ToId  uint   `json:"to_id"`
}

type rotographInfoList []rotographInfo

type GetHomeInfoResponse struct {
	Rotograph rotographInfoList `json:"rotograph"`
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
