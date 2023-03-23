package interfaces

import (
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/request"
)

type User interface {
	Create(req *request.UserCreateReq) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByPhone(phone string) (*model.User, error)
	CheckAccount(req *request.UserSignInReq) (*model.User, bool)
}

type UserManage interface {
	Delete(id uint) error
	List(pageNum, pageSize int) ([]*model.User, int64, error)
}
