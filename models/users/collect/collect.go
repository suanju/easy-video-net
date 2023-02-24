package collect

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/contribution/video"
	"Go-Live/models/users"
)

type Collect struct {
	common.PublicModel
	Uid         uint `json:"uid" gorm:"uid"`
	FavoritesID uint `json:"favorites_id" gorm:"favorites_id"`
	VideoID     uint `json:"video_id" gorm:"video_id "`

	UserInfo  users.User               `json:"userInfo" gorm:"foreignKey:Uid"`
	VideoInfo video.VideosContribution `json:"videoInfo" gorm:"foreignKey:VideoID"`
}

type CollectsList []Collect

func (Collect) TableName() string {
	return "lv_users_collect"
}

//Create 添加数据
func (c *Collect) Create() bool {
	err := global.Db.Create(c).Error
	if err != nil {
		return false
	}
	return true
}

func (l *CollectsList) FindVideoExistWhere(videoID uint) error {
	err := global.Db.Where("video_id", videoID).Find(l).Error
	return err
}
