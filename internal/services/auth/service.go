package auth

import (
	"context"
	"errors"
	authC "github.com/Rasikrr/learning_platform/internal/cache/auth"
	"github.com/Rasikrr/learning_platform/internal/clients/mail"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"

	"github.com/Rasikrr/learning_platform/internal/util"
	"math/rand/v2"
	"strings"
)

const (
	codeLength = 4
	codeChars  = "0123456789"
)

type Service interface {
	Register(ctx context.Context, email, password, passwordConfirm string) error
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
	GenerateCode(ctx context.Context) string
}

type service struct {
	mailClient mail.Client
	hasher     util.Hasher
	cache      authC.Cache
}

func NewService(mailClient mail.Client, hasher util.Hasher) Service {
	return &service{
		mailClient: mailClient,
		hasher:     hasher,
	}
}

// nolint
func (s *service) Login(_ context.Context, email, password string) (*entity.Auth, error) {
	return nil, nil
}

func (s *service) Register(ctx context.Context, email, password, passwordConfirm string) error {
	if password != passwordConfirm {
		return errors.New("passwords do not match")
	}
	code := s.GenerateCode(ctx)
	if err := s.mailClient.Send(ctx, email, []string{email}, "registration", code); err != nil {
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

func (s *service) GenerateCode(_ context.Context) string {
	nums := codeChars
	numsArr := strings.Split(nums, "")
	rand.Shuffle(len(numsArr), func(i, j int) {
		numsArr[i], numsArr[j] = numsArr[j], numsArr[i]
	})
	return strings.Join(numsArr[:codeLength], "")
}
