package service

import (
	"errors"
	"github.com/voyager-go/GoWeb/internal/interfaces"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/repository"
	"github.com/voyager-go/GoWeb/internal/request"
	"github.com/voyager-go/GoWeb/pkg/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	_ interfaces.Admin = (*AdminService)(nil)
)

type AdminService struct {
	repo      *repository.AdminRepository
	validator *validator.CustomValidator
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		repo:      repository.NewAdminRepository(db),
		validator: validator.NewCustomValidator(),
	}
}

func (a *AdminService) GetByID(id uint) (*model.Admin, error) {
	// 查询 Admin 对象，并将数据库返回结果转换为接口类型并返回
	admin, err := a.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("未查找到管理员记录")
	}
	return admin, nil
}

func (a *AdminService) SignIn(email string, password string) (string, error) {
	admin, err := a.repo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询出错并且不是记录不存在的错误，则返回错误信息
		return "", err
	}
	if admin != nil {
		err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
		if err != nil {
			return "", errors.New("邮箱与密码不匹配")
		}

	}
	return "", errors.New("未查询到指定邮箱")
}

func (s *AdminService) CheckAccount(req *request.AdminSignInReq) (*model.Admin, bool) {
	admin, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, false
	}
	return admin, true
}

func (a *AdminService) Create(req *request.AdminCreateReq) (*model.Admin, error) {
	err := a.validator.ValidateStruct(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	exist, err := a.repo.FindByNickname(req.Nickname)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询出错并且不是记录不存在的错误，则返回错误信息
		return nil, err
	}
	if exist != nil {
		// 如果昵称已经存在，则返回错误信息
		return nil, errors.New("昵称已经存在")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	newAdmin := &model.Admin{
		Nickname: req.Nickname,
		Password: string(password),
		Email:    req.Email,
	}
	err = a.repo.Create(newAdmin)
	if err != nil {
		return nil, err
	}

	// 将数据库返回结果转换为接口类型并返回
	return &model.Admin{
		ID:        newAdmin.ID,
		Nickname:  newAdmin.Nickname,
		Email:     newAdmin.Email,
		CreatedAt: newAdmin.CreatedAt,
		UpdatedAt: newAdmin.UpdatedAt,
	}, nil
}

func (a *AdminService) Update(admin *model.Admin) (*model.Admin, error) {
	// 检查管理员是否存在
	exist, err := a.repo.FindByID(admin.ID)
	if err != nil {
		return nil, err
	}
	if exist == nil {
		return nil, errors.New("该记录不存在")
	}
	// 更新 Admin 对象，并保存到数据库
	exist.Email = admin.Email
	exist.Nickname = admin.Nickname
	err = a.repo.Update(exist)
	if err != nil {
		return nil, err
	}
	// 将数据库返回结果转换为接口类型并返回
	return &model.Admin{
		ID:        exist.ID,
		Nickname:  exist.Nickname,
		Email:     exist.Email,
		CreatedAt: exist.CreatedAt,
		UpdatedAt: exist.UpdatedAt,
	}, nil
}

func (a *AdminService) List(pageNum, pageSize int) ([]*model.User, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) Delete(id uint) error {
	return a.repo.Delete(id)
}
