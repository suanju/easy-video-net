package common

import (
	"time"
)

type PublicModel struct {
	ID        uint      `json:"id" gorm:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
