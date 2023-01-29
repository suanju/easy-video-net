package video

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/contribution/video/barrage"
	"Go-Live/models/contribution/video/comments"
	"Go-Live/models/contribution/video/like"
	"Go-Live/models/users"
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

	UserInfo users.User         `json:"user_info" gorm:"foreignKey:Uid"`
	Likes    []like.Likes       `json:"likes" gorm:"foreignKey:VideoID" `
	Comments []comments.Comment `json:"comments" gorm:"foreignKey:VideoID"`
	Barrage  []barrage.Barrage  `json:"barrage" gorm:"foreignKey:VideoID"`
}

type VideosContributionList []VideosContribution

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

func (vl *VideosContributionList) GetHoneVideoList() error {
	err := global.Db.Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Order("created_at desc").Find(&vl).Error
	return err
}

//GetRecommendList 获取推荐视频
func (vl *VideosContributionList) GetRecommendList() error {
	err := global.Db.Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Order("created_at desc").Limit(7).Find(&vl).Error
	return err
}

//FindByID 根据查询
func (vc *VideosContribution) FindByID(id uint) error {
	err := global.Db.Where("id", id).Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Order("created_at desc").Find(&vc).Error
	return err
}
