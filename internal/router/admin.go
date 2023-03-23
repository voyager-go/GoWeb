package router

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/api/backend"
	"github.com/voyager-go/GoWeb/internal/middleware"
	"github.com/voyager-go/GoWeb/internal/service"
)

func initAdminRoutes(r *gin.RouterGroup, adminService *service.AdminService) {
	adminAPI := backend.NewAdminApi(adminService)
	authGroup := r.Group("").Use(middleware.BackendJWTMiddleware())
	//authGroup.GET("profile/:id", adminAPI.GetByID)
	authGroup.POST("create", adminAPI.Create)
	r.POST("sign-in", adminAPI.SignIn)
}
