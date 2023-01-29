package like

import "Go-Live/models/common"

type Likes struct {
	common.PublicModel
	Uid     uint `json:"uid" gorm:"uid"`
	VideoID uint `json:"video_id"  gorm:"video_id"`
}

func (Likes) TableName() string {
	return "lv_video_contribution_like"
}
