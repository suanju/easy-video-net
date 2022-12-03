package video

import (
	"Go-Live/models/common"
	"gorm.io/datatypes"
)

type videoContribution struct {
	common.PublicModel
	Uid       uint           `json:"uid" gorm:"uid"`
	Title     string         `json:"title" gorm:"title"`
	Cover     datatypes.JSON `json:"cover" gorm:"cover"`
	Reprinted int8           `json:"reprinted" gorm:"reprinted"`
	Timing    int8           `json:"timing" gorm:"timing"`
	Label     string         `json:"label" gorm:"label"`
	Introduce string         `json:"introduce" gorm:"introduce"`
	Heat      int            `json:"heat" gorm:"heat"`
}

func (videoContribution) TableName() string {
	return "live_video_contribution"
}
