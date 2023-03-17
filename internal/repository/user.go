package repository

import (
	"errors"
	"github.com/voyager-go/GoWeb/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&model.User{}).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	err := r.db.Model(&model.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List(page int, pageSize int) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * pageSize
	err := r.db.Model(&model.User{}).Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	err := r.db.Model(&model.User{}).Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ChangeStatus(id uint, status int) error {
	var user model.User
	if r.db.First(&user, id).RowsAffected == 0 {
		return errors.New("未查询到相关用户")
	}
	user.Status = uint8(status)
	return r.db.Save(&user).Error
}

func (r *UserRepository) DeleteByID(id uint) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error
}
