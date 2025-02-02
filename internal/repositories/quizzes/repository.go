package quizzes

import "github.com/Rasikrr/learning_platform/internal/databases"

type Repository interface {
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}
