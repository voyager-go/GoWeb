package frontend

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/config"
	"github.com/voyager-go/GoWeb/internal/model"
	"github.com/voyager-go/GoWeb/internal/request"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/helper"
	"github.com/voyager-go/GoWeb/pkg/response"
	"strconv"
)

type UserAPI struct {
	service *service.UserService
}

func NewUserAPI(service *service.UserService) *UserAPI {
	return &UserAPI{service: service}
}

// SignUp 注册用户
func (api *UserAPI) SignUp(c *gin.Context) {
	// 解析请求参数
	var req request.UserCreateReq
	if err := c.BindJSON(&req); err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}

	// 调用 UserService 的 Create 方法创建用户
	result, err := api.service.Create(&req)
	fmt.Println(result)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, err)
		return
	}

	// 返回结果给客户端
	response.OK(c, result)
}

func (api *UserAPI) SignIn(c *gin.Context) {
	var req request.UserSignInReq
	if err := c.BindJSON(&req); err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}
	user, ifExist := api.service.CheckAccount(&req)
	if !ifExist {
		response.Fail(c, response.OperationExecutionFailure, errors.New("手机与密码不匹配，请重试"))
		return
	}
	token, err := helper.GenerateToken(user.ID, config.App.Jwt.Secret, config.App.Jwt.Expired)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, errors.New("令牌生成失败，请联系管理员"))
		return
	}
	response.OK(c, token)
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
	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, response.RequestParameterError, nil)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, errors.New("类型转换时发生异常"))
		return
	}
	// 调用 UserService 的 GetByID 方法获取用户信息
	result, err := api.service.GetByID(uint(id))
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, err)
		return
	}

	// 如果用户不存在，则返回 404 错误
	if result == nil {
		response.Fail(c, response.OperationExecutionFailure, errors.New("未查找到用户"))
		return
	}

	// 返回结果给客户端
	response.OK(c, result)
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
