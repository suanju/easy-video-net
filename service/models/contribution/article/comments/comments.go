package comments

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/users"
	"easy-video-net/models/users/notice"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Comment struct {
	common.PublicModel
	Uid            uint   `json:"uid" gorm:"uid"`
	ArticleID      uint   `json:"article_id" gorm:"article_id"`
	Context        string `json:"context" gorm:"context"`
	CommentID      uint   `json:"comment_id" gorm:"comment_id"`
	CommentUserID  uint   `json:"comment_user_id" gorm:"comment_user_id"`
	CommentFirstID uint   `json:"comment_first_id" gorm:"comment_first_id"`

	UserInfo    users.User `json:"user_info" gorm:"foreignKey:Uid"`
	ArticleInfo Article    `json:"article_info" gorm:"foreignKey:ArticleID"`
}

type CommentList []Comment

func (Comment) TableName() string {
	return "lv_article_contribution_comments"
}

type Article struct {
	common.PublicModel
	Uid              uint           `json:"uid" gorm:"uid"`
	ClassificationID uint           `json:"classification_id"  gorm:"classification_id"`
	Title            string         `json:"title" gorm:"title"`
	Cover            datatypes.JSON `json:"cover" gorm:"cover"`
}

func (Article) TableName() string {
	return "lv_article_contribution"
}

//Find 根据id 查询
func (c *Comment) Find(id uint) {
	_ = global.Db.Where("id", id).Find(&c).Error
}

//Create 添加数据
func (c *Comment) Create() bool {

	err := global.Db.Transaction(func(tx *gorm.DB) error {
		articleInfo := new(Article)
		err := tx.Where("id", c.ArticleID).Find(articleInfo).Error
		if err != nil {
			return err
		}
		err = tx.Create(&c).Error
		if err != nil {
			return err
		}
		//消息通知
		if articleInfo.Uid == c.Uid {
			return nil
		}
		//添加消息通知
		ne := new(notice.Notice)
		err = ne.AddNotice(articleInfo.Uid, c.Uid, articleInfo.ID, notice.ArticleComment, c.Context)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false
	}
	return true
}

//GetCommentFirstID 获取最顶层的评论id
func (c *Comment) GetCommentFirstID() uint {
	_ = global.Db.Where("id", c.ID).Find(&c).Error
	if c.CommentID != 0 {
		c.ID = c.CommentID
		c.GetCommentFirstID()
	}
	return c.ID
}

//GetCommentUserID 获取评论id的user
func (c *Comment) GetCommentUserID() uint {
	_ = global.Db.Where("id", c.ID).Find(&c).Error
	return c.Uid
}

func (cl *CommentList) GetCommentListByIDs(ids []uint, info common.PageInfo) error {
	return global.Db.Where("article_id", ids).Preload("UserInfo").Preload("ArticleInfo").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(&cl).Error
}
