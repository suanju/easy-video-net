package like

import "Go-Live/models/common"

type Likes struct {
	common.PublicModel
	Uid            uint `json:"uid" gorm:"uid"`
	ContributionID uint `json:"contribution_id"  gorm:"contribution_id"`
}

func (Likes) TableName() string {
	return "lv_article_contribution_like"
}
