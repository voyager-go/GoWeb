package interfaces

import (
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/request"
)

type Admin interface {
	Create(req *request.AdminCreateReq) (*model.Admin, error)
	Update(admin *model.Admin) (*model.Admin, error)
	List(pageNum, pageSize int) ([]*model.User, int64, error)
	Delete(id uint) error
	SignIn(email string, password string) (string, error)
	CheckAccount(req *request.AdminSignInReq) (*model.Admin, bool)
	GetByID(id uint) (*model.Admin, error)
}
