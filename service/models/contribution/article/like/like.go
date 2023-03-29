package like

import "easy-video-net/models/common"

type Likes struct {
	common.PublicModel
	Uid       uint `json:"uid" gorm:"uid"`
	ArticleID uint `json:"article_id"  gorm:"article_id"`
}

type LikesList []Likes

func (Likes) TableName() string {
	return "lv_article_contribution_like"
}
