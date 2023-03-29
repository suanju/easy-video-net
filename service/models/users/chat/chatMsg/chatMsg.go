package chatMsg

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/users"
	"easy-video-net/models/users/chat/chatList"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Msg struct {
	common.PublicModel
	Uid     uint   `json:"uid" gorm:"uid"`
	Tid     uint   `json:"tid"  gorm:"tid"`
	Type    string `json:"type" gorm:"type"`
	Message string `json:"message" gorm:"message"`

	UInfo users.User `json:"UInfo"  gorm:"foreignKey:uid"`
	TInfo users.User `json:"TInfo"  gorm:"foreignKey:tid"`
}

type MsgList []Msg

func (Msg) TableName() string {
	return "lv_users_chat_msg"
}

func (m *Msg) AddMessage() error {
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		//添加记录
		err := tx.Create(m).Error
		if err != nil {
			return fmt.Errorf("添加聊天记录失败")
		}
		//聊天列表内添加记录
		uci := &chatList.ChatsListInfo{
			Uid:         m.Uid,
			Tid:         m.Tid,
			LastMessage: m.Message,
			LastAt:      time.Now(),
		}
		err = uci.AddChat()
		if err != nil {
			return fmt.Errorf("添加聊天列表记录失败")
		}
		tci := &chatList.ChatsListInfo{
			Uid:         m.Tid,
			Tid:         m.Uid,
			LastMessage: m.Message,
			LastAt:      time.Now(),
		}
		err = tci.AddChat()
		if err != nil {
			return fmt.Errorf("添加聊天列表记录失败")
		}
		return nil
	})
	return err
}

//FindList 获取列表
func (ml *MsgList) FindList(uid uint, tid uint) error {
	ids := make([]uint, 0)
	ids = append(ids, uid, tid)
	return global.Db.Where("uid", ids).Where("tid", ids).Preload("UInfo").Preload("TInfo").Order("created_at desc").Limit(30).Find(ml).Error
}

//GetLastMessage 获取发送的最后一条消息
func (m *Msg) GetLastMessage(uid uint, tid uint) error {
	err := global.Db.Where("uid = ? or  tid  = ? and tid = ? or tid = ? ", uid, uid, tid, tid).Order("created_at desc").Find(m).Error
	if err != nil {
		return err
	}
	return nil
}

//FindByID 根据id查询
func (m *Msg) FindByID(id uint) error {
	return global.Db.Where("id", id).Preload("UInfo").Preload("TInfo").Find(m).Error
}

//FindHistoryMsg 查询历史记录消息
func (ml *MsgList) FindHistoryMsg(uid, tid uint, lastTime time.Time) error {
	ids := make([]uint, 0)
	ids = append(ids, uid, tid)
	return global.Db.Where("uid", ids).Where("tid", ids).Where("created_at < ?", lastTime).Preload("UInfo").Preload("TInfo").Order("created_at desc").Limit(30).Find(ml).Error
}
