package repository

import (
	"errors"
	"fmt"
	"github.com/voyager-go/GoWeb/internal/costant"
	"github.com/voyager-go/GoWeb/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *model.Book) error {
	return r.db.Create(&book).Error
}

func (r *BookRepository) FindByID(id int) (*model.Book, error) {
	var book model.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(book *model.Book) error {
	return r.db.Save(&book).Error
}

func (r *BookRepository) Delete(id int) error {
	return r.db.Delete(&model.Book{}, "id = ?", id).Error
}

func (r *BookRepository) ChangeStatus(id int, newStatus int8) error {
	if newStatus != costant.StatusSerialized && newStatus != costant.StatusPublished && newStatus != costant.StatusUnpublished {
		return errors.New("状态枚举值异常")
	}
	err := r.db.Model(&model.Book{}).Where("id = ?", id).Update("status", newStatus).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("未查询到指定记录")
		}
		return err
	}
	return nil
}

func (r *BookRepository) List(title string, author string, tagStr string, page, pageSize int) ([]*model.Book, error) {
	var books []*model.Book
	q := r.db

	if title != "" {
		q.Where(fmt.Sprintf("title like %%%s%%", title))
	}
	if author != "" {
		q.Where(fmt.Sprintf("author_name like %%%s%%", author))
	}
	if tagStr != "" {
		q.Where(fmt.Sprintf("tags like %%%s%%", tagStr))
	}
	err := q.Limit(pageSize).Offset((page - 1) * pageSize).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
