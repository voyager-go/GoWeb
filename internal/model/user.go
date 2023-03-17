package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" comment:"用户ID"`
	Phone     string    `gorm:"index:phone_UNIQUE;unique;not null" comment:"手机号"`
	Status    uint8     `gorm:"not null;default:1;index:status_index" comment:"状态 1-启用 2-禁用"`
	Nickname  string    `gorm:"not null" comment:"昵称"`
	Email     string    `gorm:"index:email_UNIQUE;unique" comment:"邮箱"`
	Password  string    `gorm:"not null" comment:"密码"`
	CreatedAt time.Time `comment:"创建时间"`
	UpdatedAt time.Time `comment:"更新时间"`
}

func (u *User) TableName() string {
	return "user"
}
