package project

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(project Project) error {

	query := `
	INSERT INTO projects (id, name, api_keys, user_id) VALUES ($1, $2, $3, $4)
	`
	var pgErr *pgconn.PgError

	_, err := r.db.Exec(context.Background(), query, project.Id, project.Name, project.ApiKeys, project.UserId)

	if ok := errors.As(err, &pgErr); ok {
		if pgErr.Code == "23505" {
			return errors.New("DUPLICATE_PROJECT_NAME")
		}
	}

	return err
}