package users

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

//go:generate mockgen -destination ../mocks/users/mock.go -package mocks -source=./repository.go

type Repository interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByIDs(ctx context.Context, ids []string) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	ResetPassword(ctx context.Context, email, password string) error
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, params *entity.UpdateUserParams) error
	Delete(ctx context.Context, id string) error
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

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, deleteStmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ResetPassword(ctx context.Context, email, password string) error {
	_, err := r.db.Exec(ctx, resetPasswordStmt, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByIDStmt, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return convertModel(m)
}

func (r *repository) GetByIDs(ctx context.Context, ids []string) ([]*entity.User, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByIDsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return convertModels(mm)
}

func (r *repository) Update(ctx context.Context, params *entity.UpdateUserParams) error {
	_, err := r.db.Exec(ctx, updateStmt, params.Name, params.LastName, params.ID)
	if err != nil {
		return err
	}
	return nil
}
