package model

import (
	"github.com/voyager-go/GoWeb/pkg/formatTime"
)

// User 用户模型
type User struct {
	ID        uint            `json:"id" gorm:"primarykey" comment:"用户ID"`
	Phone     string          `json:"phone" gorm:"index:phone_UNIQUE;unique;not null" comment:"手机号"`
	Status    uint8           `json:"status" gorm:"not null;default:1;index:status_index" comment:"状态 1-启用 2-禁用"`
	Nickname  string          `json:"nickname" gorm:"not null" comment:"昵称"`
	Email     string          `json:"email" gorm:"index:email_UNIQUE;unique" comment:"邮箱"`
	Password  string          `json:"-" gorm:"not null" comment:"密码"`
	CreatedAt formatTime.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt formatTime.Time `json:"updated_at" comment:"更新时间"`
}

func (u *User) TableName() string {
	return "user"
}
