package backend

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/config"
	"github.com/voyager-go/GoWeb/internal/request"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/helper"
	"github.com/voyager-go/GoWeb/pkg/response"
)

type AdminAPI struct {
	service *service.AdminService
}

func NewAdminApi(service *service.AdminService) *AdminAPI {
	return &AdminAPI{service: service}
}

func (api *AdminAPI) SignIn(c *gin.Context) {
	var req request.AdminSignInReq
	if err := c.BindJSON(&req); err != nil {
		response.Fail(c, response.RequestParameterError, err)
		return
	}
	admin, ifExist := api.service.CheckAccount(&req)
	if !ifExist {
		response.Fail(c, response.OperationExecutionFailure, errors.New("手机与密码不匹配，请重试"))
		return
	}
	token, err := helper.GenerateToken(admin.ID, config.App.Jwt.Secret, config.App.Jwt.Expired)
	if err != nil {
		response.Fail(c, response.OperationExecutionFailure, errors.New("令牌生成失败，请联系管理员"))
		return
	}
	response.OK(c, token)
}

func (api *AdminAPI) Create(ctx *gin.Context) {
	var req request.AdminCreateReq
	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, response.RequestParameterError, err)
		return
	}
	admin, err := api.service.Create(&req)
	if err != nil {
		response.Fail(ctx, response.OperationExecutionFailure, err)
		return
	}
	response.OK(ctx, admin)
}

func (api *AdminAPI) Update(ctx *gin.Context) {

}
