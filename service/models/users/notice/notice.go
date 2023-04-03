package notice

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/users"
	"fmt"
	"gorm.io/datatypes"
)

type Notice struct {
	common.PublicModel
	Uid     uint   `json:"uid" gorm:"column:uid"`
	Cid     uint   `json:"cid" gorm:"column:cid"`
	Type    string `json:"type" gorm:"column:type"`
	ToID    uint   `json:"to_id" gorm:"column:to_id"`
	ISRead  uint   `json:"is_read" gorm:"column:is_read"`
	Content string `json:"content" gorm:"column:content"`

	VideoInfo   VideoInfo  `json:"videoInfo" gorm:"foreignKey:to_id"`
	UserInfo    users.User `json:"userinfo"  gorm:"foreignKey:cid"`
	ArticleInfo Article    `json:"articleInfo" gorm:"foreignKey:to_id"`
}

var (
	VideoComment   = "videoComment"   //视频评论
	VideoLike      = "videoLike"      //视频点赞
	ArticleComment = "articleComment" //文章评论
	ArticleLike    = "articleLike"    //文章点赞

)

type NoticesList []Notice

func (Notice) TableName() string {
	return "lv_users_notice"
}

//VideoInfo 临时加一个video模型解决依赖循环
type VideoInfo struct {
	common.PublicModel
	Uid   uint           `json:"uid" gorm:"uid"`
	Title string         `json:"title" gorm:"title"`
	Video datatypes.JSON `json:"video" gorm:"video"`
	Cover datatypes.JSON `json:"cover" gorm:"cover"`
}

func (VideoInfo) TableName() string {
	return "lv_video_contribution"
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

func (nl *NoticesList) GetNoticeList(info common.PageInfo, messageType []string, uid uint) error {
	if len(messageType) > 0 {
		return global.Db.Where("uid", uid).Where("type", messageType).Preload("VideoInfo").Preload("ArticleInfo").Preload("UserInfo").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(nl).Error
	} else {
		return global.Db.Where("uid", uid).Preload("VideoInfo").Preload("ArticleInfo").Preload("UserInfo").Limit(info.Size).Offset((info.Page - 1) * info.Size).Order("created_at desc").Find(nl).Error

	}
}

func (n *Notice) AddNotice(uid uint, cid uint, tid uint, tp string, content string) error {
	n.Uid = uid
	n.Cid = cid
	n.ToID = tid
	n.Type = tp
	n.Content = content
	n.ISRead = 0
	return global.Db.Create(n).Error
}
func (n *Notice) Delete(uid uint, cid uint, tid uint, tp string) error {
	return global.Db.Where(&Notice{
		Uid:  uid,
		Cid:  cid,
		Type: tp,
		ToID: tid,
	}).Delete(n).Error
}

func (n *Notice) GetUnreadNum(uid uint) *int64 {
	num := new(int64)
	err := global.Db.Model(n).Where("uid", uid).Where("is_read", 0).Count(num).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return num
}

func (n *Notice) ReadAll(uid uint) error {
	return global.Db.Where(Notice{Uid: uid, ISRead: 0}).Updates(&Notice{ISRead: 1}).Error
}
