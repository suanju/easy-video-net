package video

import (
	"Go-Live/models/common"
	"time"
)

type CreateVideoContributionReceiveStruct struct {
	Video           string    `json:"video" binding:"required"`
	VideoUploadType string    `json:"videoUploadType"  binding:"required"`
	Cover           string    `json:"cover" binding:"required"`
	CoverUploadType string    `json:"coverUploadType" binding:"required"`
	Title           string    `json:"title" binding:"required"`
	Reprinted       *bool     `json:"reprinted" binding:"required"`
	Timing          *bool     `json:"timing" binding:"required"`
	TimingTime      time.Time `json:"timingTime"`
	Label           []string  `json:"label" binding:"required"`
	Introduce       string    `json:"introduce" binding:"required"`
	VideoDuration   int64     `json:"videoDuration" binding:"required"`
}

type GetVideoContributionByIDReceiveStruct struct {
	VideoID uint `json:"video_id"`
}

type SendVideoBarrageReceiveStruct struct {
	Author string  `json:"author"`
	Color  uint    `json:"color"`
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Time   float64 `json:"time"`
	Type   uint    `json:"type"`
	Token  string  `json:"token"`
}

type GetVideoBarrageReceiveStruct struct {
	ID string `json:"id"`
}

type GetVideoBarrageListReceiveStruct struct {
	ID string `json:"id"`
}

type VideosPostCommentReceiveStruct struct {
	VideoID   uint   `json:"video_id"`
	Content   string `json:"content"`
	ContentID uint   `json:"content_id"`
}

type GetVideoCommentReceiveStruct struct {
	PageInfo common.PageInfo `json:"pageInfo"`
	VideoID  uint            `json:"video_id" binding:"required"`
}
