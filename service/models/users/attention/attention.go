package attention

import (
	"easy-video-net/global"
	"easy-video-net/models/common"
	"easy-video-net/models/users"
)

type Attention struct {
	common.PublicModel
	Uid         uint `json:"uid" gorm:"uid"`
	AttentionID uint `json:"attention_id" gorm:"attention_id"`

	UserInfo          users.User `json:"user_info" gorm:"foreignKey:Uid"`
	AttentionUserInfo users.User `json:"attention_user_info" gorm:"foreignKey:AttentionID"`
}

type AttentionsList []Attention

func (Attention) TableName() string {
	return "lv_users_attention"
}

//Create 添加数据
func (at *Attention) Create() bool {
	err := global.Db.Create(&at).Error
	if err != nil {
		return false
	}
	return true
}

//Delete 删除数据
func (at *Attention) Delete() bool {
	err := global.Db.Where("uid", at.Uid).Updates(&at).Error
	if err != nil {
		return false
	}
	return true
}

//Attention 关注用户
func (at *Attention) Attention(uid uint, aid uint) bool {
	err := global.Db.Where(Attention{Uid: uid, AttentionID: aid}).Find(&at).Error
	if at.ID > 0 {
		//已关注
		err = global.Db.Where("id", at.ID).Delete(&at).Error
	} else {
		//未关注
		err = global.Db.Create(&Attention{Uid: uid, AttentionID: aid}).Error
	}

	if err != nil {
		return false
	}
	return true
}

//IsAttention  判断是否关注用户
func (at *Attention) IsAttention(uid uint, aid uint) bool {
	_ = global.Db.Where(Attention{Uid: uid, AttentionID: aid}).Find(&at).Error
	if at.ID > 0 {
		return true
	} else {
		return false
	}
}

//GetAttentionNum 获取关注数量
func (at *Attention) GetAttentionNum(uid uint) (*int64, error) {
	num := new(int64)
	err := global.Db.Model(at).Where(Attention{Uid: uid}).Count(num).Error
	if err != nil {
		return num, err
	}
	return num, nil
}

//GetVermicelliNum 获取粉丝数量
func (at *Attention) GetVermicelliNum(uid uint) (*int64, error) {
	num := new(int64)
	err := global.Db.Model(at).Where(Attention{AttentionID: uid}).Count(num).Error
	if err != nil {
		return num, err
	}
	return num, nil
}

//GetAttentionList 获取关注列表
func (al *AttentionsList) GetAttentionList(uid uint) error {
	err := global.Db.Where("uid", uid).Preload("AttentionUserInfo").Find(al).Error
	if err != nil {
		return err
	}
	return nil
}

//GetVermicelliList 获取粉丝列表
func (al *AttentionsList) GetVermicelliList(uid uint) error {
	err := global.Db.Where("attention_id", uid).Preload("UserInfo").Find(al).Error
	if err != nil {
		return err
	}
	return nil
}

//GetAttentionListByIdArr 获取关注列表 id数组
func (al *AttentionsList) GetAttentionListByIdArr(uid uint) (arr []uint, err error) {
	arr = make([]uint, 0)
	err = global.Db.Where("uid", uid).Find(al).Error
	if err != nil {
		return arr, err
	}
	//自需要数组
	for _, v := range *al {
		arr = append(arr, v.AttentionID)
	}
	return arr, nil
}
