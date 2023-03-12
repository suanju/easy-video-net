package chatList

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
)

type ChatsListInfo struct {
	common.PublicModel
	Uid         uint   `json:"uid" gorm:"uid"`
	Tid         uint   `json:"tid"  gorm:"tid"`
	LastMessage string `json:"last_message" gorm:"last_message"`

	ToUserInfo users.User `json:"toUserInfo"  gorm:"foreignKey:tid"`
}

type ChatList []ChatsListInfo

func (ChatsListInfo) TableName() string {
	return "lv_users_chat_list"
}

//AddChat 添加记录
func (i *ChatsListInfo) AddChat() error {
	//判断是否存在
	is := &ChatsListInfo{}
	err := global.Db.Where("uid = ? And tid = ?", i.Uid, i.Tid).Find(is).Error
	if err != nil {
		return err
	}
	if is.ID != 0 {
		return nil
	} else {
		return global.Db.Create(i).Error
	}
}

//DeleteChat 删除记录
func (i *ChatsListInfo) DeleteChat(tid uint, uid uint) error {
	return global.Db.Where("uid = ? And tid = ?", uid, tid).Delete(i).Error
}

//GetListByIO 获取聊天记录
func (cl *ChatList) GetListByIO(uid uint) error {
	return global.Db.Where("uid", uid).Preload("ToUserInfo").Order("updated_at desc").Find(cl).Error
}
