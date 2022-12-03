package video

import "gorm.io/datatypes"

type CreateVideoContributionReceiveStruct struct {
	Video          string         `json:"video" binding:"required"`
	VideoInterface string         `json:"videoInterface" binding:"required"`
	Cover          string         `json:"cover" binding:"required"`
	CoverInterface string         `json:"coverInterface" binding:"required"`
	Title          string         `json:"title" binding:"required"`
	Reprinted      *bool          `json:"reprinted" binding:"required"`
	Timing         *bool          `json:"timing" binding:"required"`
	TimingTime     datatypes.Time `json:"timingTime"`
	Label          string         `json:"label" binding:"required"`
	Introduce      string         `json:"introduce" binding:"required"`
}
