package record

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/article"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/users"
	"fmt"
)

type Record struct {
	common.PublicModel
	Uid  uint   `json:"column:uid"`
	Type string `json:"type" gorm:"column:type"`
	ToId uint   `json:"to_id" gorm:"column:to_id"`

	VideoInfo   video.VideosContribution     `json:"videoInfo" gorm:"foreignKey:to_id"`
	Userinfo    users.User                   `json:"users.User"  gorm:"foreignKey:uid"`
	ArticleInfo article.ArticlesContribution `json:"articleInfo" gorm:"foreignKey:to_id"`
}

type RecordsList []Record

func (Record) TableName() string {
	return "lv_users_record"
}

//ClearRecord 清空历史记录
func (r *Record) ClearRecord(uid uint) error {
	return global.Db.Where("uid", uid).Delete(r).Error
}

//DeleteRecordByID 删除历史记录根据id
func (r *Record) DeleteRecordByID(id uint, uid uint) error {
	err := global.Db.Where("id", id).Find(r).Error
	if err != nil {
		return err
	}
	if r.Uid != uid {
		return fmt.Errorf("非法操作")
	}
	return global.Db.Where("id", id).Delete(r).Error
}

func (r *Record) AddVideoRecord(uid uint, to uint) error {
	err := global.Db.Where(Record{Uid: uid, Type: "video", ToId: to}).Find(r).Error
	if err != nil {
		return err
	}
	if r.ID <= 0 {
		//创建记录
		r.Uid = uid
		r.Type = "video"
		r.ToId = to
		return global.Db.Create(r).Error
	} else {
		//更新记录
		return global.Db.Where("id", r.ID).Updates(r).Error
	}
}
func (r *Record) AddArticleRecord(uid uint, to uint) error {
	err := global.Db.Where(Record{Uid: uid, Type: "article", ToId: to}).Find(r).Error
	if err != nil {
		return err
	}
	if r.ID <= 0 {
		//创建记录
		r.Uid = uid
		r.Type = "article"
		r.ToId = to
		return global.Db.Create(r).Error
	} else {
		//更新记录
		return global.Db.Where("id", r.ID).Updates(r).Error
	}
}

func (r *Record) AddLiveRecord(uid uint, to uint) error {
	err := global.Db.Where(Record{Uid: uid, Type: "live", ToId: to}).Find(r).Error
	if err != nil {
		return err
	}
	if r.ID <= 0 {
		//创建记录
		r.Uid = uid
		r.Type = "live"
		r.ToId = to
		return global.Db.Create(r).Error
	} else {
		//更新记录
		return global.Db.Where("id", r.ID).Updates(r).Error
	}
}

func (l *RecordsList) GetRecordListByUid(uid uint, info common.PageInfo) error {
	err := global.Db.Where("uid", uid).Preload("VideoInfo.UserInfo").Preload("ArticleInfo.UserInfo").Preload("Userinfo.LiveInfo").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(&l).Error
	return err
}
