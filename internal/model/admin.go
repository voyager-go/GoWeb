package model

import (
	"github.com/voyager-go/GoWeb/pkg/formatTime"
)

// Admin 管理员表
type Admin struct {
	ID        uint            `gorm:"primarykey" json:"id" comment:"管理员ID"`
	Nickname  string          `gorm:"not null;uniqueIndex:idx_nickname" json:"nickname" comment:"用户名"  validate:"required,min=2,max=60"`
	Password  string          `gorm:"not null" json:"-" comment:"密码" validate:"required,min=4,max=20"`
	Email     string          `gorm:"not null;uniqueIndex:idx_email" json:"email" comment:"邮箱" validate:"required,email"`
	CreatedAt formatTime.Time `gorm:"not null" json:"created_at" comment:"创建时间"`
	UpdatedAt formatTime.Time `gorm:"not null" json:"updated_at" comment:"更新时间"`
}

func (a *Admin) TableName() string {
	return "admin"
}
