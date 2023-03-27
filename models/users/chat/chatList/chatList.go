package chatList

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
	"fmt"
	"time"
)

type ChatsListInfo struct {
	common.PublicModel
	Uid         uint      `json:"uid" gorm:"uid"`
	Tid         uint      `json:"tid"  gorm:"tid"`
	Unread      int       `json:"unread" gorm:"unread"`
	LastMessage string    `json:"last_message" gorm:"last_message"`
	LastAt      time.Time `json:"last_at" gorm:"last_at"`

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
		//存在即更新
		global.Db.Model(is).Updates(map[string]interface{}{"last_at": i.LastAt, "last_message": i.LastMessage})
		return nil
	} else {
		return global.Db.Create(i).Error
	}
}

//DeleteChat 删除记录
func (i *ChatsListInfo) DeleteChat(tid uint, uid uint) error {
	return global.Db.Where("uid = ? and tid = ?", uid, tid).Delete(i).Error
}

//GetListByIO 获取聊天记录
func (cl *ChatList) GetListByIO(uid uint) error {
	return global.Db.Where("uid", uid).Preload("ToUserInfo").Order("updated_at desc").Find(cl).Error
}

func (i *ChatsListInfo) UnreadAutocorrection(uid uint, tid uint) error {
	err := global.Db.Where(ChatsListInfo{Uid: uid, Tid: tid}).Find(i).Error
	if err != nil {
		return err
	}
	if i.ID > 0 {
		i.Unread++
		return global.Db.Save(i).Error
	}
	return fmt.Errorf("无记录")
}

func (i *ChatsListInfo) UnreadEmpty(uid uint, tid uint) error {
	err := global.Db.Where(ChatsListInfo{Uid: uid, Tid: tid}).Find(i).Error
	if err != nil {
		return err
	}
	if i.ID > 0 {
		i.Unread = 0
		return global.Db.Save(i).Error
	}
	return fmt.Errorf("情况失败")
}

func (i *ChatsListInfo) FindByID(uid uint, tid uint) error {
	return global.Db.Where(ChatsListInfo{Uid: uid, Tid: tid}).Find(i).Error
}

//GetUnreadNumber 获取该用户未读消息数量
func (i ChatsListInfo) GetUnreadNumber(uid uint) *int64 {
	num := new(int64)
	err := global.Db.Model(i).Select("SUM(unread) as total_unread").Where(ChatsListInfo{Uid: uid}).Scan(num).Error
	if err != nil {
		fmt.Println(err)
	}
	return num
}
