package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// JWTMiddleware 基于JWT的鉴权中间件
func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid authorization token",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			// 验证通过，将用户ID存储到上下文中
			c.Set("phone", claims.Phone)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization token",
		})
		c.Abort()
		return
	}
}

// GenerateToken 生成JWT Token
func GenerateToken(phone string, secret string, expireTime time.Duration) (string, error) {
	expirationTime := time.Now().Add(expireTime)
	claims := &Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
