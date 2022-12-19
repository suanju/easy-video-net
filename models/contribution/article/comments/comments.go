package comments

import "Go-Live/models/common"

type Comment struct {
	common.PublicModel
	Uid            uint   `json:"uid" gorm:"uid"`
	ContributionID uint   `json:"contribution_Id" gorm:"contribution_Id"`
	Context        string `json:"context" gorm:"context"`
	CommentID      string `json:"comment_Id" gorm:"comment_Id"`
}

func (Comment) TableName() string {
	return "live_article_contribution_comments"
}
