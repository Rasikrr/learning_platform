package auth

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go
type registerRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type confirmRegisterRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func convertToModel(auth *entity.Auth) Auth {
	return Auth{
		RefreshToken: auth.RefreshToken,
		AccessToken:  auth.AccessToken,
	}
}
