package auth

import (
	"context"
	"errors"
	authC "github.com/Rasikrr/learning_platform/internal/cache/auth"
	"github.com/Rasikrr/learning_platform/internal/clients/mail"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	usersR "github.com/Rasikrr/learning_platform/internal/repositories/users"
	"github.com/golang-jwt/jwt"
	"time"

	"github.com/Rasikrr/learning_platform/internal/util"
	"math/rand/v2"
	"strings"
)

const (
	codeLength           = 4
	codeChars            = "0123456789"
	accessTokenLifeTime  = time.Minute * 15
	refreshTokenLifeTime = time.Hour * 24
	secret               = "rasik1234"
)

type Service interface {
	Register(ctx context.Context, email, password, passwordConfirm string) error
	ConfirmRegister(ctx context.Context, email, code string) (*entity.Auth, error)
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
	GenerateCode(ctx context.Context) string
}

type service struct {
	mailClient mail.Client
	repository usersR.Repository
	hasher     util.Hasher
	cache      authC.Cache
}

func NewService(
	mailClient mail.Client,
	usersRepo usersR.Repository,
	hasher util.Hasher,
	cache authC.Cache,
) Service {
	return &service{
		mailClient: mailClient,
		hasher:     hasher,
		repository: usersRepo,
		cache:      cache,
	}
}

// nolint
func (s *service) Login(ctx context.Context, email, password string) (*entity.Auth, error) {
	user, err := s.repository.GetByEmail(ctx, email)
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
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user already exists")
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
	if err := s.repository.Create(ctx, user); err != nil {
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

func (s *service) generateTokens(_ context.Context, user *entity.User) (*entity.Auth, error) {
	claims := entity.NewClaims(user, accessTokenLifeTime, false)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	claims = entity.NewClaims(user, refreshTokenLifeTime, true)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &entity.Auth{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
