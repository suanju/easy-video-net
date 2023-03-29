package common

import (
	"time"
)

type PublicModel struct {
	ID        uint      `json:"id" gorm:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

type Img struct {
	Src string `json:"src" gorm:"src"`
	Tp  string `json:"type" gorm:"type"`
}
