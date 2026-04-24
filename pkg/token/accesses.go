package token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MaximTretjakov/nofelet-web/config"
)

type JWTToken struct {
	config *config.Config
}

// NewJWTToken - Создание новой структуры анонимного токена
func NewJWTToken(config *config.Config) *JWTToken {
	return &JWTToken{
		config: config,
	}
}

// CreateToken - Создание анонимного токена
func (t *JWTToken) CreateToken(ctx context.Context, login string) (string, error) {
	token, err := t.create(ctx, login)
	return token, err
}

// create - Создание анонимного токена
func (t *JWTToken) create(ctx context.Context, login string) (string, error) {
	claims := jwt.MapClaims{
		"sub": login,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtWithClaims.SignedString([]byte(t.config.JWT.ValidationKey))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", t.config.JWT.Prefix, token), nil
}
