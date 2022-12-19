package video

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"gorm.io/datatypes"
	"time"
)

type VideosContribution struct {
	common.PublicModel
	Uid        uint           `json:"uid" gorm:"uid"`
	Title      string         `json:"title" gorm:"title"`
	Video      datatypes.JSON `json:"video" gorm:"video"`
	Cover      datatypes.JSON `json:"cover" gorm:"cover"`
	Reprinted  int8           `json:"reprinted" gorm:"reprinted"`
	Timing     int8           `json:"timing" gorm:"timing"`
	TimingTime time.Time      `json:"timingTime"  gorm:"timing_Time"`
	Label      string         `json:"label" gorm:"label"`
	Introduce  string         `json:"introduce" gorm:"introduce"`
	Heat       int            `json:"heat" gorm:"heat"`
}

func (VideosContribution) TableName() string {
	return "live_video_contribution"
}

//Create 添加数据
func (vc *VideosContribution) Create() bool {
	err := global.Db.Create(&vc).Error
	if err != nil {
		return false
	}
	return true
}
