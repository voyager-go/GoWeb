package backend

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/service"
	"math"
	"strconv"
)

type UserAPI struct {
	service *service.UserManageService
}

func NewUserAPI(service *service.UserManageService) *UserAPI {
	return &UserAPI{service: service}
}

// Delete 删除用户
func (api *UserAPI) Delete(c *gin.Context) {
	// 获取用户 ID 参数
	id, err := GetIDParam(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用 UserService 的 Delete 方法删除用户
	err = api.service.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回结果给客户端
	c.Status(204)
}

// List 分页查询用户列表
func (api *UserAPI) List(c *gin.Context) {
	// 获取分页参数
	pageNum, pageSize, err := GetPaginationParams(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用 UserService 的 List 方法查询用户列表
	users, total, err := api.service.List(pageNum, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 构造分页结果
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := &model.Pagination{
		PageNum:    pageNum,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		List:       users,
	}

	// 返回结果给客户端
	c.JSON(200, pagination)
}

// GetIDParam 从 URL 参数中获取用户 ID
func GetIDParam(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return uint(id), nil
}

// GetPaginationParams 从 URL 参数中获取分页参数
func GetPaginationParams(c *gin.Context) (int, int, error) {
	pageNumStr := c.DefaultQuery("page_num", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")

	fmt.Println(pageNumStr, pageSizeStr)
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		return 0, 0, errors.New("invalid pageNum parameter")
	}
	if pageNum < 1 {
		return 0, 0, errors.New("pageNum must be greater than or equal to 1")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, 0, errors.New("invalid pageSize parameter")
	}
	if pageSize < 1 || pageSize > 100 {
		return 0, 0, errors.New("pageSize must be between 1 and 100")
	}

	return pageNum, pageSize, nil
}
