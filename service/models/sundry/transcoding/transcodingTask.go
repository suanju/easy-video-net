package transcodingTask

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
)

var Aliyun = "aliyun"

type TranscodingTask struct {
	common.PublicModel
	TaskID     string `json:"interface"  gorm:"column:task_id"`
	VideoID    uint   `json:"video_id"  gorm:"column:video_id"`
	Resolution int    `json:"resolution"  gorm:"column:resolution"`
	Dst        string `json:"dst" gorm:"column:dst"`
	Status     int    `json:"method"  gorm:"column:status"`
	Type       string `json:"path" gorm:"column:type"`
}

func (TranscodingTask) TableName() string {
	return "lv_transcoding_task"
}

func (t *TranscodingTask) AddTask() bool {
	err := global.Db.Create(t).Error
	if err != nil {
		return false
	}
	return true
}
func (t *TranscodingTask) Save() bool {
	err := global.Db.Save(t).Error
	if err != nil {
		return false
	}
	return true
}
func (t *TranscodingTask) GetInfoByTaskID(id string) error {
	return global.Db.Where("task_id", id).Find(t).Error
}
