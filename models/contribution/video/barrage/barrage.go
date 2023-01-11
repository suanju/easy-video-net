package barrage

import "Go-Live/models/common"

type Barrage struct {
	common.PublicModel
	Uid      uint   `json:"uid" gorm:"uid"`
	VideID   uint   `json:"video_id" gorm:"video_id"`
	Text     string `json:"text" gorm:"text"`
	Color    string `json:"color" gorm:"color"`
	Position string `json:"position" gorm:"position"`
}
