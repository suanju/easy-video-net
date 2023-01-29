package comments

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
)

type Comment struct {
	common.PublicModel
	Uid            uint   `json:"uid" gorm:"uid"`
	VideoID        uint   `json:"video_id" gorm:"contribution_id"`
	Context        string `json:"context" gorm:"context"`
	CommentID      uint   `json:"comment_id" gorm:"comment_id"`
	CommentUserID  uint   `json:"comment_user_id" gorm:"comment_user_id"`
	CommentFirstID uint   `json:"comment_first_id" gorm:"comment_first_id"`

	UserInfo users.User `json:"user_info" gorm:"foreignKey:Uid"`
}

func (Comment) TableName() string {
	return "lv_video_contribution_comments"
}

//Find 根据id 查询
func (c *Comment) Find(id uint) {
	_ = global.Db.Where("id", id).Find(&c).Error
}

//Create 添加数据
func (c *Comment) Create() bool {
	err := global.Db.Create(&c).Error
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
