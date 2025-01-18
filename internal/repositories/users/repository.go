package users

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entities"
	"sync"
	"time"
)

//go:generate mockgen -destination ../mocks/users/mock.go -package mocks -source=./repository.go

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context) ([]*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, email string) error
	UpdateName(ctx context.Context, email, name string) error
}

type repository struct {
	mut *sync.RWMutex
	db  map[string]*model
}

func NewRepository() Repository {
	return &repository{
		mut: &sync.RWMutex{},
		db:  make(map[string]*model, 10),
	}
}

func (r *repository) GetByEmail(_ context.Context, email string) (*entities.User, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	m, ok := r.db[email]
	if !ok {
		return nil, errors.New("user with this email not found")
	}
	return convertModel(m), nil
}

func (r *repository) GetAll(_ context.Context) ([]*entities.User, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	out := make(models, 0, len(r.db))
	for _, v := range r.db {
		out = append(out, v)
	}
	return convertModels(out), nil
}

func (r *repository) Create(_ context.Context, user *entities.User) error {
	r.mut.Lock()
	defer r.mut.Unlock()
	m := convertToModel(user)
	r.db[m.Email] = m
	return nil
}

func (r *repository) Delete(_ context.Context, email string) error {
	r.mut.Lock()
	defer r.mut.Unlock()

	if _, ok := r.db[email]; !ok {
		return errors.New("user with this email not found")
	}
	delete(r.db, email)
	return nil
}

func (r *repository) UpdateName(_ context.Context, email, name string) error {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.db[email].Name = name
	r.db[email].UpdatedAt = time.Now()
	return nil
}
