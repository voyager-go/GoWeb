package router

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/api/frontend"
	"github.com/voyager-go/GoWeb/internal/middleware"
	"github.com/voyager-go/GoWeb/internal/service"
)

func initUserRoutes(r *gin.RouterGroup, userService *service.UserService) {
	userAPI := frontend.NewUserAPI(userService)
	authGroup := r.Group("").Use(middleware.JWTMiddleware())
	authGroup.GET("profile/:id", userAPI.GetByID)
	r.POST("sign-up", userAPI.SignUp)
	r.POST("sign-in", userAPI.SignIn)
}
