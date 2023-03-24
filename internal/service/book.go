package service

import (
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/repository"
	"github.com/voyager-go/GoWeb/internal/request"
	"github.com/voyager-go/GoWeb/pkg/validator"
	"gorm.io/gorm"
)

type BookService struct {
	repo      *repository.BookRepository
	validator *validator.CustomValidator
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		repo:      repository.NewBookRepository(db),
		validator: validator.NewCustomValidator(), // 初始化验证器
	}
}

// GetBookByID 根据ID获取书籍信息
func (s *BookService) GetBookByID(id int) (*model.Book, error) {
	return s.repo.FindByID(id)
}

// Create 添加新书籍
func (s *BookService) Create(book *model.Book) error {
	return s.repo.Create(book)
}

// Update 更新已有书籍
func (s *BookService) Update(book *model.Book) error {
	return s.repo.Update(book)
}

// Delete 根据ID删除书籍
func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

// ChangeStatus 更改书籍状态
func (s *BookService) ChangeStatus(id int, newStatus int8) error {
	return s.repo.ChangeStatus(id, newStatus)
}

// Count 获取所有书籍的数量
func (s *BookService) Count() (int64, error) {
	return s.repo.Count()
}

// List 根据标签获取书籍列表
func (s *BookService) List(req *request.BookListReq) ([]*model.Book, int64, error) {
	var books []*model.Book
	books, err := s.repo.List(req.Title, req.Author, req.Tags, req.Page, req.PageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return books, total, nil
}
