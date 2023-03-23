package router

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/orm"
)

func InitRouter(srv *gin.Engine) {
	// 初始化后端路由
	initBackendRouter(srv)
	// 初始化前端路由
	initFrontendRouter(srv)
}

func initBackendRouter(srv *gin.Engine) {
	adminService := service.NewAdminService(orm.Conn)
	// 后端分组路由
	api := srv.Group("/backend")
	initAdminRoutes(api.Group("admin"), adminService)
}

func initFrontendRouter(srv *gin.Engine) {
	userService := service.NewUserService(orm.Conn)
	// 前端分组路由
	api := srv.Group("/api")
	// 初始化各个分组的路由
	initUserRoutes(api.Group("user"), userService)
}
