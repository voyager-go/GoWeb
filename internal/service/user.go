package service

import (
	"errors"
	"github.com/voyager-go/GoWeb/internal/costant"
	"github.com/voyager-go/GoWeb/internal/interfaces"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/repository"
	"github.com/voyager-go/GoWeb/internal/request"
	"github.com/voyager-go/GoWeb/pkg/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	_ interfaces.User       = (*UserService)(nil)
	_ interfaces.UserManage = (*UserManageService)(nil)
)

//ipPackage := qqwry.NewQQwry(config.Cfg.Server.QqwryPath)
//ipPackage.Find(req.PublishedIp)

type UserService struct {
	repo      *repository.UserRepository
	validator *validator.CustomValidator
}

type UserManageService struct {
	repo      *repository.UserRepository
	validator *validator.CustomValidator
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo:      repository.NewUserRepository(db),
		validator: validator.NewCustomValidator(), // 初始化验证器
	}
}

func NewUserManageService(db *gorm.DB) *UserManageService {
	return &UserManageService{
		repo:      repository.NewUserRepository(db),
		validator: validator.NewCustomValidator(),
	}
}

func (s *UserService) CheckAccount(req *request.UserSignInReq) (*model.User, bool) {
	user, err := s.repo.FindByPhone(req.Phone)
	if err != nil {
		return nil, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, false
	}
	return user, true
}

func (s *UserService) Create(req *request.UserCreateReq) (*model.User, error) {
	// 验证用户参数
	err := s.validator.ValidateStruct(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	// 检查手机号是否已经存在
	exist, err := s.repo.FindByPhone(req.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询出错并且不是记录不存在的错误，则返回错误信息
		return nil, err
	}
	if exist != nil {
		// 如果手机号已经存在，则返回错误信息
		return nil, errors.New("手机号已经存在")
	}

	// 生成密码哈希

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	if req.Status == 0 {
		req.Status = uint8(costant.UserEnabled)
	}
	// 构造新的 User 对象，并插入到数据库中
	newUser := &model.User{
		Phone:    req.Phone,
		Status:   req.Status,
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: string(password),
	}
	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	// 将数据库返回结果转换为接口类型并返回
	return &model.User{
		ID:        newUser.ID,
		Phone:     newUser.Phone,
		Status:    newUser.Status,
		Nickname:  newUser.Nickname,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (s *UserService) Update(user *model.User) (*model.User, error) {
	// 检查用户是否存在
	exist, err := s.repo.FindByID(user.ID)
	if err != nil {
		return nil, err
	}
	if exist == nil {
		return nil, errors.New("该记录不存在")
	}

	// 更新 User 对象，并保存到数据库中
	exist.Phone = user.Phone
	exist.Status = user.Status
	exist.Nickname = user.Nickname
	exist.Email = user.Email
	err = s.repo.Update(exist)
	if err != nil {
		return nil, err
	}

	// 将数据库返回结果转换为接口类型并返回
	return &model.User{
		ID:        exist.ID,
		Phone:     exist.Phone,
		Status:    exist.Status,
		Nickname:  exist.Nickname,
		Email:     exist.Email,
		CreatedAt: exist.CreatedAt,
		UpdatedAt: exist.UpdatedAt,
	}, nil
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	// 查询 User 对象，并将数据库返回结果转换为接口类型并返回
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &model.User{
		ID:        user.ID,
		Phone:     user.Phone,
		Status:    user.Status,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) GetByPhone(phone string) (*model.User, error) {
	// 查询 User 对象，并将数据库返回结果转换为接口类型并返回
	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &model.User{
		ID:        user.ID,
		Phone:     user.Phone,
		Status:    user.Status,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserManageService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserManageService) List(pageNum, pageSize int) ([]*model.User, int64, error) {
	// 计算分页参数
	offset := (pageNum - 1) * pageSize

	// 查询用户总数
	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	// 查询用户分页列表
	users, err := s.repo.List(offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 将数据库返回结果转换为接口类型并返回
	result := make([]*model.User, len(users))
	for i, u := range users {
		result[i] = &model.User{
			ID:        u.ID,
			Phone:     u.Phone,
			Status:    u.Status,
			Nickname:  u.Nickname,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
	}

	// 返回结果
	return result, total, nil
}
