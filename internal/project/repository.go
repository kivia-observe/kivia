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

func (r Repository) FindById(id string) (Project, error) {

	var project Project

	query := `
	SELECT id, name, api_keys, user_id, created_at FROM projects WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&project.Id, &project.Name, &project.ApiKeys, &project.UserId, &project.CreatedAt)

	return project, err
}

func (r Repository) ExistsById(id string) (bool, error) {

	var exists bool

	query := `
	SELECT EXISTS(SELECT 1 FROM projects WHERE id = $1)
	`

	row := r.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&exists)

	return exists, err
}