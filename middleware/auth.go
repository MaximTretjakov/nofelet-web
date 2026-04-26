package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

var (
	ErrMissingAuthHeader  = errors.New("missing authorization header")
	ErrInvalidTokenFormat = errors.New("invalid token format")
	ErrTokenExpired       = errors.New("invalid or expired token")
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newError(ErrMissingAuthHeader))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newError(ErrInvalidTokenFormat))
			return
		}

		tokenString := parts[1]

		// 3. Парсим и валидируем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверка алгоритма подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		// Если токен протух (exp) или подпись неверна — Valid будет false
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newError(ErrTokenExpired))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["user_id"])
		}

		c.Next()
	}
}

func newError(err error) view.SimpleErrorBody {
	return view.SimpleErrorBody{
		Data: struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		},
	}
}
