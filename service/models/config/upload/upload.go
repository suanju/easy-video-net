package upload

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
)

type Upload struct {
	common.PublicModel
	Interfaces string  `json:"interface"  gorm:"column:interface"`
	Method     string  `json:"method"  gorm:"column:method"`
	Path       string  `json:"path" gorm:"column:path"`
	Quality    float64 `json:"quality"  gorm:"column:quality"`
}

func (Upload) TableName() string {
	return "lv_upload_method"
}

//IsExistByField 根据字段判断用户是否存在
func (um *Upload) IsExistByField(field string, value any) bool {
	err := global.Db.Where(field, value).Find(&um).Error
	if err != nil {
		return false
	}
	if um.ID <= 0 {
		return false
	}
	return true
}
