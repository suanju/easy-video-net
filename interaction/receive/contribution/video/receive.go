package video

import (
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
