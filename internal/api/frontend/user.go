package frontend

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/response"
	"strconv"
)

type UserAPI struct {
	service *service.UserService
}

func NewUserAPI(service *service.UserService) *UserAPI {
	return &UserAPI{service: service}
}

// Register 注册用户
func (api *UserAPI) Register(c *gin.Context) {
	// 解析请求参数
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}

	// 调用 UserService 的 Create 方法创建用户
	result, err := api.service.Create(&user)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, err)
		return
	}

	// 返回结果给客户端
	response.OK(c, result)
}

// Update 更新用户信息
func (api *UserAPI) Update(c *gin.Context) {
	// 解析请求参数
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	// 调用 UserService 的 Update 方法更新用户信息
	result, err := api.service.Update(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回结果给客户端
	c.JSON(200, result)
}

// GetByID 根据 ID 获取用户信息
func (api *UserAPI) GetByID(c *gin.Context) {
	// 获取用户 ID 参数
	id, err := GetIDParam(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用 UserService 的 GetByID 方法获取用户信息
	result, err := api.service.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 如果用户不存在，则返回 404 错误
	if result == nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	// 返回结果给客户端
	c.JSON(200, result)
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
