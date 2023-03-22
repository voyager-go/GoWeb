package router

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/api/frontend"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/orm"
)

func InitRouter(srv *gin.Engine) {
	userService := service.NewUserService(orm.Conn)

	// 分组路由
	api := srv.Group("/api")

	// 初始化各个分组的路由
	initUserRoutes(api.Group("/user"), userService)
}

func initUserRoutes(r *gin.RouterGroup, userService *service.UserService) {
	userAPI := frontend.NewUserAPI(userService)
	r.POST("sign-up", userAPI.SignUp)
	r.GET("profile", userAPI.GetByID)
}
