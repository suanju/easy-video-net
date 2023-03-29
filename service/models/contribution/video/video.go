package video

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/video/barrage"
	"easy-video-net/models/contribution/video/comments"
	"easy-video-net/models/contribution/video/like"
	"easy-video-net/models/users"
	"gorm.io/datatypes"
	"gorm.io/gorm"
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

	UserInfo users.User           `json:"user_info" gorm:"foreignKey:Uid"`
	Likes    like.LikesList       `json:"likes" gorm:"foreignKey:VideoID" `
	Comments comments.CommentList `json:"comments" gorm:"foreignKey:VideoID"`
	Barrage  barrage.BarragesList `json:"barrage" gorm:"foreignKey:VideoID"`
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

//Delete 删除数据
func (vc *VideosContribution) Delete(id uint, uid uint) bool {
	if global.Db.Where("id", id).Find(&vc).Error != nil {
		return false
	}
	if vc.Uid != uid {
		return false
	}
	if global.Db.Delete(&vc).Error != nil {
		return false
	}
	return true
}

//Update 更新数据
func (vc *VideosContribution) Update(info map[string]interface{}) bool {
	err := global.Db.Model(vc).Updates(info).Error
	if err != nil {
		return false
	}
	return true
}

//FindByID 根据查询
func (vc *VideosContribution) FindByID(id uint) error {
	return global.Db.Where("id", id).Preload("Likes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserInfo").Order("created_at desc")
	}).Preload("Barrage").Preload("UserInfo").Order("created_at desc").Find(&vc).Error
}

//GetVideoComments 获取评论
func (vc *VideosContribution) GetVideoComments(ID uint, info common.PageInfo) bool {
	err := global.Db.Where("id", ID).Preload("Likes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserInfo").Order("created_at desc").Limit(info.Size).Offset((info.Page - 1) * info.Size)
	}).Find(vc).Error
	if err != nil {
		return false
	}
	return true
}

//Watch 添加播放
func (vc *VideosContribution) Watch(id uint) error {
	return global.Db.Model(vc).Where("id", id).Updates(map[string]interface{}{"heat": gorm.Expr("Heat  + ?", 1)}).Error
}

//GetVideoListBySpace 获取个人空间视频列表
func (vl *VideosContributionList) GetVideoListBySpace(id uint) error {
	return global.Db.Where("uid", id).Preload("Likes").Preload("Comments").Preload("Barrage").Order("created_at desc").Find(&vl).Error
}

//GetDiscussVideoCommentList 获取个人发布的视频和评论信息
func (vl *VideosContributionList) GetDiscussVideoCommentList(id uint) error {
	return global.Db.Where("uid", id).Preload("Comments").Find(&vl).Error
}

func (vl *VideosContributionList) GetVideoManagementList(info common.PageInfo, uid uint) error {
	return global.Db.Where("uid", uid).Preload("Likes").Preload("Comments").Preload("Barrage").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(&vl).Error
}
func (vl *VideosContributionList) GetHoneVideoList(info common.PageInfo) error {
	//首页加载13个铺满后续15个
	var offset int
	if info.Page == 1 {
		info.Size = 11
		offset = (info.Page - 1) * info.Size
	}
	offset = (info.Page-2)*info.Size + 11

	return global.Db.Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Limit(info.Size).Offset(offset).Order("created_at desc").Find(&vl).Error
}

//GetRecommendList 获取推荐视频
func (vl *VideosContributionList) GetRecommendList() error {
	return global.Db.Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Order("created_at desc").Limit(7).Find(&vl).Error
}

func (vl *VideosContributionList) Search(info common.PageInfo) error {
	return global.Db.Where("`title` LIKE ?", "%"+info.Keyword+"%").Preload("Likes").Preload("Comments").Preload("Barrage").Preload("UserInfo").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(&vl).Error

}
