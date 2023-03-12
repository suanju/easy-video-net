package chatMsg

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
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

//FindList 获取列表
func (ml *MsgList) FindList(uid uint, tid uint) error {
	ids := make([]uint, 0)
	ids = append(ids, uid, tid)
	return global.Db.Where("uid", ids).Where("tid", ids).Preload("UInfo").Preload("TInfo").Order("created_at desc").Find(ml).Error
}

//GetLastMessage 获取发送的最后一条消息
func (m *Msg) GetLastMessage(uid uint, tid uint) error {
	err := global.Db.Where("uid = ? or  tid  = ? and tid = ? or tid = ? ", uid, uid, tid, tid).Order("created_at desc").Find(m).Error
	if err != nil {
		return err
	}
	return nil
}
