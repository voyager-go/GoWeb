package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/GoWeb/internal/config"
	"github.com/voyager-go/GoWeb/internal/service"
	"github.com/voyager-go/GoWeb/pkg/helper"
	"github.com/voyager-go/GoWeb/pkg/orm"
	"github.com/voyager-go/GoWeb/pkg/response"
	"strings"
)

// BackendJWTMiddleware 基于JWT的鉴权中间件
func BackendJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.App.Jwt.Secret
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, response.AuthorizationHeaderIsRequired, nil)
			c.Abort()
			return
		}

		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := jwt.ParseWithClaims(tokenString, &helper.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.Fail(c, response.InvalidAuthorizationToken, err)
				c.Abort()
				return
			}
			response.Fail(c, response.BadRequest, err)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*helper.Claims); ok && token.Valid {
			// 验证该ID是否对应数据库中的一个用户
			admin, err := service.NewAdminService(orm.Conn).GetByID(claims.ID)
			if err != nil {
				response.Fail(c, response.InvalidAuthorizationToken, nil)
				c.Abort()
				return
			}
			// 验证通过，将用户ID存储到上下文中
			c.Set("admin", admin)
			c.Next()
			return
		}

		response.Fail(c, response.InvalidAuthorizationToken, nil)
		c.Abort()
		return
	}
}
