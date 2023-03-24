package model

import (
	"github.com/voyager-go/GoWeb/pkg/formatTime"
)

type Book struct {
	ID                int             `gorm:"primary_key;comment:'书籍ID'"`
	Title             string          `gorm:"not null;comment:'书籍名称';index"`
	Description       string          `gorm:"not null;comment:'书籍简介'"`
	Tags              string          `gorm:"comment:'书籍标签';index"`
	Cover             string          `gorm:"not null;comment:'书籍封面'"`
	AuthorName        string          `gorm:"not null;comment:'作者名称';index"`
	AuthorDescription string          `gorm:"comment:'作者简介'"`
	Status            int8            `gorm:"not null;default:1;comment:'书籍状态 1-连载中 2-已发布 3-已下架'"`
	CreatedAt         formatTime.Time `gorm:"comment:'创建时间'"`
	UpdatedAt         formatTime.Time `gorm:"comment:'更新时间'"`
}

func (b *Book) TableName() string {
	return "book_info"
}
