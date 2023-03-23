package repository

import (
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/pkg/formatTime"
	"gorm.io/gorm"
	"time"
)

// NewAdminRepository 新建管理员仓库
func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db}
}

type AdminRepository struct {
	db *gorm.DB
}

// FindByID 根据 ID 查询管理员
func (r *AdminRepository) FindByID(id uint) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// FindByNickname 根据用户名查询管理员
func (r *AdminRepository) FindByNickname(nickname string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.Where("nickname = ?", nickname).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// FindByEmail 根据邮箱查询管理员
func (r *AdminRepository) FindByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// Create 创建管理员
func (r *AdminRepository) Create(admin *model.Admin) error {
	now := formatTime.Time{Time: time.Now()}
	admin.CreatedAt = now
	admin.UpdatedAt = now
	return r.db.Create(admin).Error
}

// Update 更新管理员
func (r *AdminRepository) Update(admin *model.Admin) error {
	admin.UpdatedAt = formatTime.Time{Time: time.Now()}
	return r.db.Save(admin).Error
}

// Delete 删除管理员
func (r *AdminRepository) Delete(id uint) error {
	return r.db.Delete(&model.Admin{}, id).Error
}
