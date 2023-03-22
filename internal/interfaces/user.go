package interfaces

import (
	"github.com/voyager-go/GoWeb/internal/model"
)

type User interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByPhone(phone string) (*model.User, error)
}

type UserManage interface {
	Delete(id uint) error
	List(pageNum, pageSize int) ([]*model.User, int64, error)
}
