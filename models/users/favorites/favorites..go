package favorites

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
	"Go-Live/models/users/collect"
	"fmt"
	"gorm.io/datatypes"
)

type Favorites struct {
	common.PublicModel
	Uid     uint           `json:"uid" gorm:"uid"`
	Title   string         `json:"title" gorm:"title"`
	Content string         `json:"content" gorm:"content"`
	Cover   datatypes.JSON `json:"cover" gorm:"type:json;comment:cover"`
	Max     int            `json:"max" gorm:"max"`

	UserInfo    users.User           `json:"userInfo" gorm:"foreignKey:Uid"`
	CollectList collect.CollectsList `json:"collectList"  gorm:"foreignKey:FavoritesID"`
}

type FavoriteList []Favorites

func (Favorites) TableName() string {
	return "lv_users_favorites"
}

//Find 查询
func (f *Favorites) Find(id uint) bool {
	err := global.Db.Where("id", id).Preload("CollectList").Order("created_at desc").Find(&f).Error
	if err != nil {
		return false
	}
	return true
}

//Create 添加数据
func (f *Favorites) Create() bool {
	err := global.Db.Create(&f).Error
	if err != nil {
		return false
	}
	return true
}

//AloneTitleCreate 单标题创建
func (f *Favorites) AloneTitleCreate() bool {
	err := global.Db.Create(&f).Error
	if err != nil {
		return false
	}
	return true
}

//Update 更新数据
func (f *Favorites) Update() bool {
	err := global.Db.Updates(&f).Error
	if err != nil {
		return false
	}
	return true
}

//Delete 删除数据
func (f *Favorites) Delete(id uint, uid uint) error {
	err := global.Db.Where("id", id).Find(&f).Error
	if err != nil {
		return fmt.Errorf("查询失败")
	}
	if f.ID <= 0 {
		return fmt.Errorf("收藏夹不存在")
	}
	err = global.Db.Delete(&f).Error
	if f.Uid != uid {
		return fmt.Errorf("非创建者不可删除")
	}
	//删除收藏记录
	cl := new(collect.Collect)
	if !cl.DetectByFavoritesID(id) {
		return fmt.Errorf("删除收藏记录失败")
	}
	if err != nil {
		return fmt.Errorf("删除失败")
	}
	return nil
}

func (fl *FavoriteList) GetFavoritesList(id uint) error {
	err := global.Db.Where("uid", id).Preload("UserInfo").Preload("CollectList").Order("created_at desc").Find(fl).Error
	if err != nil {
		return err
	}
	return nil
}
