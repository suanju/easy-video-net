package video

import (
	"easy-video-net/models/common"
)

type CreateVideoContributionReceiveStruct struct {
	Video           string   `json:"video" binding:"required"`
	VideoUploadType string   `json:"videoUploadType" binding:"required"`
	Cover           string   `json:"cover" binding:"required"`
	CoverUploadType string   `json:"coverUploadType" binding:"required"`
	Title           string   `json:"title" binding:"required"`
	Reprinted       *bool    `json:"reprinted" binding:"required"`
	Label           []string `json:"label"`
	Introduce       string   `json:"introduce" binding:"required"`
	VideoDuration   int64    `json:"videoDuration" binding:"required"`
	Media           *string  `json:"media" binding:"required"`
}

type UpdateVideoContributionReceiveStruct struct {
	ID              uint     `json:"id" binding:"required"`
	Cover           string   `json:"cover" binding:"required"`
	CoverUploadType string   `json:"coverUploadType" binding:"required"`
	Title           string   `json:"title" binding:"required"`
	Reprinted       *bool    `json:"reprinted" binding:"required"`
	Label           []string `json:"label"`
	Introduce       string   `json:"introduce" binding:"required"`
}

type GetVideoContributionByIDReceiveStruct struct {
	VideoID uint `json:"video_id"`
}

type SendVideoBarrageReceiveStruct struct {
	Author string  `json:"author"`
	Color  uint    `json:"color" binding:"required"`
	ID     string  `json:"id" binding:"required"`
	Text   string  `json:"text" binding:"required"`
	Time   float64 `json:"time"`
	Type   uint    `json:"type"`
	Token  string  `json:"token" binding:"required"`
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

type GetVideoManagementListReceiveStruct struct {
	PageInfo common.PageInfo `json:"page_info"`
}

type DeleteVideoByIDReceiveStruct struct {
	ID uint `json:"id"`
}

type LikeVideoReceiveStruct struct {
	VideoID uint `json:"video_id"`
}
