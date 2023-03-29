package rotograph

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"gorm.io/datatypes"
)

type Rotograph struct {
	common.PublicModel
	Title string         `json:"title" gorm:"title"`
	Cover datatypes.JSON `json:"cover" gorm:"cover"`
	Color string         `json:"color" gorm:"color" `
	Type  string         `json:"type" gorm:"type"`
	ToId  uint           `json:"to_id" gorm:"to_id"`
}

type List []Rotograph

func (Rotograph) TableName() string {
	return "lv_hone_rotograph"
}

func (l *List) GetAll() error {
	err := global.Db.Find(&l).Error
	return err
}
