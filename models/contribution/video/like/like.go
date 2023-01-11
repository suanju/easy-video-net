package like

import "Go-Live/models/common"

type Likes struct {
	common.PublicModel
	Uid    uint `json:"uid" gorm:"uid"`
	VideID uint `json:"video_id"  gorm:"video_id"`
}

func (Likes) TableName() string {
	return "lv_article_contribution_like"
}
