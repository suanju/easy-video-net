package rotograph

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"gorm.io/datatypes"
)

type Rotograph struct {
	common.PublicModel
	Title string         `json:"title" gorm:"column:title"`
	Cover datatypes.JSON `json:"cover" gorm:"column:cover"`
	Color string         `json:"color" gorm:"column:color" `
	Type  string         `json:"type" gorm:"column:type"`
	ToId  uint           `json:"to_id" gorm:"column:to_id"`
}

type List []Rotograph

func (Rotograph) TableName() string {
	return "lv_hone_rotograph"
}

func (l *List) GetAll() error {
	err := global.Db.Find(&l).Error
	return err
}
