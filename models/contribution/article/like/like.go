package like

import "Go-Live/models/common"

type Likes struct {
	common.PublicModel
	Uid            uint `json:"uid" gorm:"uid"`
	ContributionID uint `json:"contribution_Id"`
}

func (Likes) TableName() string {
	return "live_article_contribution_like"
}
