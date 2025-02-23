package auth

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go

type ResetPasswordRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ConfirmResetPasswordRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ConfirmRegisterRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func convertToModel(auth *entity.Auth) Auth {
	return Auth{
		RefreshToken: auth.RefreshToken,
		AccessToken:  auth.AccessToken,
	}
}
