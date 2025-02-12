package auth

import (
	"context"
	"errors"
	authC "github.com/Rasikrr/learning_platform/internal/cache/auth"
	"github.com/Rasikrr/learning_platform/internal/clients/mail"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	usersR "github.com/Rasikrr/learning_platform/internal/repositories/users"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
	"unicode"

	"github.com/Rasikrr/learning_platform/internal/util"
	"math/rand/v2"
	"strings"
)

const (
	codeLength        = 4
	codeChars         = "0123456789"
	passwordMinLength = 8
)

type Service interface {
	Register(ctx context.Context, email, password, passwordConfirm string) error
	ConfirmRegister(ctx context.Context, email, code string) (*entity.Auth, error)
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
	ResetPassword(ctx context.Context, email, password, passwordConfirm string) error
	ConfirmResetPassword(ctx context.Context, email, code string) error
	GenerateCode(ctx context.Context) string
	CheckToken(ctx context.Context, token string) (*entity.Session, error)
	RefreshToken(ctx context.Context, token string) (*entity.Auth, error)
}

type service struct {
	secret               string
	accessTokenLifeTime  time.Duration
	refreshTokenLifeTime time.Duration

	mailClient      mail.Client
	usersRepository usersR.Repository
	hasher          util.Hasher
	cache           authC.Cache
}

func NewService(
	secret string,
	accessTokenLifeTime time.Duration,
	refreshTokenLifeTime time.Duration,
	mailClient mail.Client,
	usersRepo usersR.Repository,
	hasher util.Hasher,
	cache authC.Cache,
) Service {
	return &service{
		secret:               secret,
		accessTokenLifeTime:  accessTokenLifeTime,
		refreshTokenLifeTime: refreshTokenLifeTime,
		mailClient:           mailClient,
		hasher:               hasher,
		usersRepository:      usersRepo,
		cache:                cache,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (*entity.Auth, error) {
	user, err := s.usersRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err := s.hasher.CheckPassword(user.Password, password); err != nil {
		return nil, err
	}
	return s.generateTokens(ctx, user)
}

func (s *service) Register(ctx context.Context, email, password, passwordConfirm string) error {
	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}
	user, err := s.usersRepository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user already exists")
	}

	if err := s.validatePassword(ctx, password); err != nil {
		return err
	}

	code := s.GenerateCode(ctx)
	if err := s.mailClient.Send(ctx, []string{email}, "registration", code); err != nil {
		return err
	}
	if err := s.cache.SetCode(ctx, email, code); err != nil {
		return err
	}
	passHash, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}
	if err := s.cache.SetPasswordHash(ctx, email, passHash); err != nil {
		return err
	}
	return nil
}

func (s *service) ResetPassword(ctx context.Context, email, password, passwordConfirm string) error {
	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}
	user, err := s.usersRepository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	err = s.validatePassword(ctx, password)
	if err != nil {
		return err
	}
	code := s.GenerateCode(ctx)
	err = s.mailClient.Send(ctx, []string{email}, "reset_password", code)
	if err != nil {
		return err
	}
	if err := s.cache.SetCode(ctx, email, code); err != nil {
		return err
	}
	passHash, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}
	return s.cache.SetPasswordHash(ctx, email, passHash)
}

func (s *service) ConfirmResetPassword(ctx context.Context, email, code string) error {
	codeFromCache, err := s.cache.GetCode(ctx, email)
	if err != nil {
		return err
	}
	if codeFromCache != code {
		return errors.New("code is invalid")
	}
	passHash, err := s.cache.GetPasswordHash(ctx, email)
	if err != nil {
		return err
	}
	if err = s.usersRepository.ResetPassword(ctx, email, passHash); err != nil {
		return err
	}
	return nil
}

func (s *service) ConfirmRegister(ctx context.Context, email, code string) (*entity.Auth, error) {
	codeFromCache, err := s.cache.GetCode(ctx, email)
	if err != nil {
		return nil, err
	}
	if codeFromCache != code {
		return nil, errors.New("code is invalid")
	}
	passHash, err := s.cache.GetPasswordHash(ctx, email)
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(email, passHash)
	if err := s.usersRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return s.generateTokens(ctx, user)
}

func (s *service) GenerateCode(_ context.Context) string {
	nums := codeChars
	numsArr := strings.Split(nums, "")
	rand.Shuffle(len(numsArr), func(i, j int) {
		numsArr[i], numsArr[j] = numsArr[j], numsArr[i]
	})
	return strings.Join(numsArr[:codeLength], "")
}

func (s *service) CheckToken(_ context.Context, token string) (*entity.Session, error) {
	claims, isRefresh, err := s.parseJWT(token)
	if err != nil {
		return nil, err
	}
	if isRefresh {
		return nil, errors.New("access token expected")
	}
	return getSession(claims)
}

func (s *service) RefreshToken(ctx context.Context, token string) (*entity.Auth, error) {
	claims, isRefresh, err := s.parseJWT(token)
	if err != nil {
		return nil, err
	}
	if !isRefresh {
		return nil, errors.New("refresh token expected")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email is empty")
	}
	user, err := s.usersRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return s.generateTokens(ctx, user)
}

func (s *service) generateTokens(_ context.Context, user *entity.User) (*entity.Auth, error) {
	claims := entity.NewClaims(user, s.accessTokenLifeTime, false)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	claims = entity.NewClaims(user, s.refreshTokenLifeTime, true)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	return &entity.Auth{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *service) parseJWT(tokenString string) (jwt.MapClaims, bool, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	tokenString = strings.TrimPrefix(tokenString, "bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, false, err
	}
	claims := token.Claims.(jwt.MapClaims)
	isRefresh, ok := claims["is_refresh"].(bool)
	if !ok {
		return nil, false, errors.New("is_refresh is empty")
	}
	return claims, isRefresh, nil
}

func (s *service) validatePassword(_ context.Context, password string) error {
	if len(password) < passwordMinLength {
		return errors.New("password is too short")
	}
	var (
		hasUppercase, hasLowercase, hasDigit, hasSpecial bool
	)
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUppercase = true
		case unicode.IsLower(r):
			hasLowercase = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}
	if !hasUppercase || !hasLowercase || !hasDigit || !hasSpecial {
		return errors.New("password must contain at least one uppercase, lowercase, digit and special character")
	}
	return nil
}

func getSession(claims jwt.MapClaims) (*entity.Session, error) {
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email is empty")
	}
	roleString, ok := claims["account_role"].(string)
	if !ok {
		return nil, errors.New("account_role is empty")
	}
	role, err := enum.AccountRoleString(roleString)
	if err != nil {
		return nil, err
	}
	uuidString, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("user_id is empty")
	}
	userID, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}
	return &entity.Session{
		UserID: userID,
		Email:  email,
		Role:   role,
	}, nil
}
