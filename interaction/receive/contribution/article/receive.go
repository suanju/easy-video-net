package article

import (
	"Go-Live/models/common"
	"time"
)

type CreateArticleContributionReceiveStruct struct {
	Cover                         string    `json:"cover" binding:"required"`
	CoverUploadType               string    `json:"coverUploadType" binding:"required"`
	ArticleContributionUploadType string    `json:"articleContributionUploadType" binding:"required"`
	Title                         string    `json:"title" binding:"required"`
	Timing                        *bool     `json:"timing" binding:"required"`
	TimingTime                    time.Time `json:"timingTime"`
	Label                         []string  `json:"label" binding:"required"`
	Content                       string    `json:"content" binding:"required"`
	Comments                      *bool     `json:"comments"  binding:"required"`
	ClassificationID              uint      `json:"classification_id"`
}

type GetArticleContributionListByUserReceiveStruct struct {
	UserID uint `json:"userID" binding:"required"`
}

type GetArticleContributionByIDReceiveStruct struct {
	ArticleID uint `json:"articleID" binding:"required"`
}
type ArticlesPostCommentReceiveStruct struct {
	ArticleID uint   `json:"article_id"`
	Content   string `json:"content"`
	ContentID uint   `json:"content_id"`
}

type GetArticleCommentReceiveStruct struct {
	PageInfo  common.PageInfo `json:"pageInfo"`
	ArticleID uint            `json:"articleID" binding:"required"`
}
