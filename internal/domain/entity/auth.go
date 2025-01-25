package entity

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Auth struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	AccountRole string `json:"account_role"`
	IsRefresh   bool   `json:"is_refresh"`
	jwt.StandardClaims
}

func NewClaims(user *User, lifetime time.Duration, isRefresh bool) *Claims {
	if isRefresh {
		return &Claims{
			Email:     user.Email,
			IsRefresh: isRefresh,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(lifetime).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
	}
	return &Claims{
		UserID:      user.ID.String(),
		Email:       user.Email,
		AccountRole: user.AccountRole.String(),
		IsRefresh:   isRefresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
