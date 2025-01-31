package users

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination ../mocks/users/mock.go -package mocks -source=./repository.go

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, email string) error
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByEmailStmt, email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return convertModel(m)
}

func (r *repository) Create(ctx context.Context, user *entity.User) error {
	m := convertToModel(user)
	_, err := r.db.Exec(ctx, createStmt, m.ID, m.Name, m.LastName, m.Email, m.Password, m.AccountRole)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, email string) error {
	_, err := r.db.Exec(ctx, deleteStmt, email)
	if err != nil {
		return err
	}
	return nil
}
