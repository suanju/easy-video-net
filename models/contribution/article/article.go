package article

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/contribution/article/classification"
	"Go-Live/models/contribution/article/comments"
	"Go-Live/models/contribution/article/like"
	"Go-Live/models/users"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ArticlesContribution struct {
	common.PublicModel
	Uid                uint           `json:"uid" gorm:"uid"`
	ClassificationID   uint           `json:"classification_id"  gorm:"classification_id"`
	Title              string         `json:"title" gorm:"title"`
	Cover              datatypes.JSON `json:"cover" gorm:"cover"`
	Timing             int8           `json:"timing" gorm:"timing"`
	TimingTime         time.Time      `json:"timingTime"  gorm:"timing_Time"`
	Label              string         `json:"label" gorm:"label"`
	Content            string         `json:"content" gorm:"content"`
	ContentStorageType string         `json:"content_Storage_Type" gorm:"content_storage_type"`
	IsComments         int8           `json:"is_comments" gorm:"is_comments"`
	Heat               int            `json:"heat" gorm:"heat"`

	//光联表

	UserInfo       users.User                    `json:"user_info" gorm:"foreignKey:Uid"`
	Likes          []like.Likes                  `json:"likes" gorm:"foreignKey:ContributionID" `
	Comments       []comments.Comment            `json:"comments" gorm:"foreignKey:ContributionID"`
	Classification classification.Classification `json:"classification"  gorm:"foreignKey:ClassificationID"`
}

type ArticlesContributionList []ArticlesContribution

func (ArticlesContribution) TableName() string {
	return "lv_article_contribution"
}

//Create 添加数据
func (vc *ArticlesContribution) Create() bool {
	err := global.Db.Create(&vc).Error
	if err != nil {
		return false
	}
	return true
}

//GetList 查询数据类型
func (l *ArticlesContributionList) GetList(info common.PageInfo) bool {
	err := global.Db.Preload("Likes").Preload("Classification").Preload("Comments").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(l).Error
	if err != nil {
		return false
	}
	return true
}

func (vc *ArticlesContribution) Update(info map[string]interface{}) bool {
	err := global.Db.Model(vc).Updates(info).Error
	if err != nil {
		return false
	}
	return true
}

//GetListByUid 查询单个用户
func (l *ArticlesContributionList) GetListByUid(uid uint) bool {
	err := global.Db.Where("uid", uid).Preload("Likes").Preload("Classification").Preload("Comments").Order("created_at desc").Find(l).Error
	if err != nil {
		return false
	}
	return true
}

//GetAllCount 查询所有文章数量
func (l *ArticlesContributionList) GetAllCount(cu *int64) bool {
	err := global.Db.Find(l).Count(cu).Error
	if err != nil {
		return false
	}
	return true
}

//GetInfoByID 查询单个文章
func (vc *ArticlesContribution) GetInfoByID(ID uint) bool {
	err := global.Db.Where("id", ID).Preload("Likes").Preload("Classification").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserInfo").Order("created_at desc")
	}).Find(vc).Error
	if err != nil {
		return false
	}
	return true
}

//GetArticleComments 获取评论
func (vc *ArticlesContribution) GetArticleComments(ID uint, info common.PageInfo) bool {
	err := global.Db.Where("id", ID).Preload("Likes").Preload("Classification").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("UserInfo").Order("created_at desc").Limit(info.Size).Offset((info.Page - 1) * info.Size)
	}).Find(vc).Error
	if err != nil {
		return false
	}
	return true
}

//GetArticleBySpace 获取个人空间专栏列表
func (l *ArticlesContributionList) GetArticleBySpace(id uint) error {
	err := global.Db.Where("uid", id).Preload("Likes").Preload("Classification").Preload("Comments").Order("created_at desc").Find(l).Error
	if err != nil {
		return err
	}
	return nil
}

//GetArticleManagementList 创作空间获取个人发布专栏
func (l *ArticlesContributionList) GetArticleManagementList(info common.PageInfo, id uint) error {
	err := global.Db.Where("uid", id).Preload("Likes").Preload("Classification").Preload("Comments").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(l).Error
	if err != nil {
		return err
	}
	return nil
}

func (vc *ArticlesContribution) Delete(id uint, uid uint) bool {
	err := global.Db.Where("id", id).Find(vc).Error
	if err != nil {
		return false
	}
	if vc.Uid != uid {
		return false
	}
	err = global.Db.Delete(vc).Error
	if err != nil {
		return false
	}
	return true
}
