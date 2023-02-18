package attention

import (
	"Go-Live/global"
	"Go-Live/models/common"
)

type Attention struct {
	common.PublicModel
	Uid         uint `json:"uid" gorm:"uid"`
	AttentionID uint `json:"attention_id" gorm:"attention_id"`
}

func (Attention) TableName() string {
	return "lv_users_attention"
}

//Create 添加数据
func (li *Attention) Create() bool {
	err := global.Db.Create(&li).Error
	if err != nil {
		return false
	}
	return true
}

//Delete 删除数据
func (li *Attention) Delete() bool {
	err := global.Db.Where("uid", li.Uid).Updates(&li).Error
	if err != nil {
		return false
	}
	return true
}

//Attention 关注用户
func (li *Attention) Attention(uid uint, aid uint) bool {
	err := global.Db.Where(Attention{Uid: uid, AttentionID: aid}).Find(&li).Error
	if li.ID > 0 {
		//已关注
		err = global.Db.Where("id", li.ID).Delete(&li).Error
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
func (li *Attention) IsAttention(uid uint, aid uint) bool {
	_ = global.Db.Where(Attention{Uid: uid, AttentionID: aid}).Find(&li).Error
	if li.ID > 0 {
		return true
	} else {
		return false
	}
}

//GetAttentionNum 获取关注数量
func (li *Attention) GetAttentionNum(uid uint) (*int64, error) {
	num := new(int64)
	err := global.Db.Model(li).Where(Attention{Uid: uid}).Count(num).Error
	if err != nil {
		return num, err
	}
	return num, nil
}

//GetVermicelliNum 获取粉丝数量
func (li *Attention) GetVermicelliNum(uid uint) (*int64, error) {
	num := new(int64)
	err := global.Db.Model(li).Where(Attention{AttentionID: uid}).Count(num).Error
	if err != nil {
		return num, err
	}
	return num, nil
}
