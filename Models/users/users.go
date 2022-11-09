package users

import (
	"Go-Live/Global"
	"Go-Live/Models/common"
	"crypto/md5"
	"fmt"
	"time"
)

//User 表结构体
type User struct {
	common.PublicModel
	Email     string    `json:"email" gorm:"email"`
	Username  string    `json:"username" gorm:"username"`
	Openid    string    `json:"openid" gorm:"openid"`
	Salt      string    `json:"salt" gorm:"salt"`
	Password  string    `json:"password" gorm:"password"`
	Photo     string    `json:"photo" gorm:"photo"`
	Gender    int       `json:"gender" gorm:"gender"`
	BirthDate time.Time `json:"birth_date" gorm:"birth_date"`
	IsVisible int       `json:"is_visible" gorm:"is_visible"`
	Signature string    `json:"signature" gorm:"signature"`
}

func (User) TableName() string {
	return "lv_users"
}

//Update 更新数据
func (us *User) Update(user *User) bool {
	err := Global.Db.Where("id", us.ID).Updates(&user).Error
	if err != nil {
		return false
	}
	return true
}

//Create 添加数据
func (us *User) Create() bool {
	err := Global.Db.Create(&us).Error
	if err != nil {
		return false
	}
	return true
}

//IsExistByField 根据字段判断用户是否存在
func (us *User) IsExistByField(field string, value any) bool {
	err := Global.Db.Where(field, value).Find(&us).Error
	if err != nil {
		return false
	}
	if us.ID <= 0 {
		return false
	}
	return true
}

//IfPasswordCorrect 根据字段判断用户是否存在
func (us *User) IfPasswordCorrect(password string) bool {
	passwordImport := fmt.Sprintf("%s%s%s", us.Salt, password, us.Salt)
	passwordImport = fmt.Sprintf("%x", md5.Sum([]byte(passwordImport)))
	if passwordImport != us.Password {
		return false
	}
	return true
}
