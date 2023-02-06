package barrage

import (
	"Go-Live/global"
	"Go-Live/models/common"
)

type Barrage struct {
	common.PublicModel
	Uid     uint    `json:"uid" gorm:"uid"`
	VideoID uint    `json:"video_id" gorm:"video_id"`
	Time    float64 `json:"time" gorm:"time"`
	Author  string  `json:"author" gorm:"author"`
	Type    uint    `json:"type" gorm:"type"`
	Text    string  `json:"text" gorm:"text"`
	Color   uint    `json:"color" gorm:"color"`
}

type BarragesList []Barrage

func (Barrage) TableName() string {
	return "lv_video_contribution_barrage"
}

func (b *Barrage) Create() bool {
	err := global.Db.Create(&b).Error
	if err != nil {
		return false
	}
	return true
}

//GetVideoBarrageByID 查询视频弹幕
func (l *BarragesList) GetVideoBarrageByID(id uint) bool {
	err := global.Db.Where("video_id", id).Find(&l).Error
	if err != nil {
		return false
	}
	return true
}
