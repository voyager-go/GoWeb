package backend

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/helper"
	"github.com/voyager-go/GoWeb/pkg/response"
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}

	// 调用 UserService 的 Delete 方法删除用户
	err = api.service.Delete(uint(id))
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, errors.New("删除用户时发生异常，请重试"))
		return
	}

	// 返回结果给客户端
	response.OK(c, nil)
}

// List 分页查询用户列表
func (api *UserAPI) List(c *gin.Context) {
	// 获取分页参数
	pageNum, pageSize, err := helper.ParsePagination(c)
	if err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}

	// 调用 UserService 的 List 方法查询用户列表
	users, total, err := api.service.List(pageNum, pageSize)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, err)
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
	response.OK(c, pagination)
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
