package article

import (
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
}

type GetArticleContributionListByUserReceiveStruct struct {
	UserID uint `json:"userID" binding:"required"`
}

type GetArticleContributionByIDReceiveStruct struct {
	ArticleID uint `json:"articleID" binding:"required"`
}
