package video

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/contribution/video/comments"
	"Go-Live/models/contribution/video/like"
	"gorm.io/datatypes"
	"time"
)

type VideosContribution struct {
	common.PublicModel
	Uid           uint           `json:"uid" gorm:"uid"`
	Title         string         `json:"title" gorm:"title"`
	Video         datatypes.JSON `json:"video" gorm:"video"`
	Cover         datatypes.JSON `json:"cover" gorm:"cover"`
	VideoDuration int64          `json:"video_duration" gorm:"video_duration"`
	Reprinted     int8           `json:"reprinted" gorm:"reprinted"`
	Timing        int8           `json:"timing" gorm:"timing"`
	TimingTime    time.Time      `json:"timingTime"  gorm:"timing_Time"`
	Label         string         `json:"label" gorm:"label"`
	Introduce     string         `json:"introduce" gorm:"introduce"`
	Heat          int            `json:"heat" gorm:"heat"`

	Likes    []like.Likes       `json:"likes" gorm:"foreignKey:VideID" `
	Comments []comments.Comment `json:"comments" gorm:"foreignKey:VideID"`
}

func (VideosContribution) TableName() string {
	return "lv_video_contribution"
}

//Create 添加数据
func (vc *VideosContribution) Create() bool {
	err := global.Db.Create(&vc).Error
	if err != nil {
		return false
	}
	return true
}
func (vc *VideosContribution) GetHoneVideoList() error {
	err := global.Db.Find(&vc).Error
	return err
}
