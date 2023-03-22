package repository

import (
	"github.com/voyager-go/GoWeb/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"phone":      user.Phone,
		"status":     user.Status,
		"nickname":   user.Nickname,
		"email":      user.Email,
		"password":   user.Password,
		"updated_at": user.UpdatedAt,
	}).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	user := new(model.User)
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByPhone(phone string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("phone = ?", phone).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List(offset, limit int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
