package apikey

import (
	"context"

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

func (r Repository) Save(apiKey ApiKey) error {

	query := `
		INSERT INTO api_keys (id, name, key, user_id, project_id, revoked) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(context.Background(), query, apiKey.Id, apiKey.Name, apiKey.Key, apiKey.UserId, apiKey.ProjectId, apiKey.Revoked)

	return err
}

func (r Repository) FindByKey(hashedKey string) (ApiKey, error) {

	var apiKey ApiKey 

		query := `
		SELECT id, name, key, user_id, project_id, revoked, created_at FROM api_keys WHERE key = $1
	`

	row := r.db.QueryRow(context.Background(), query, hashedKey)

	err := row.Scan(&apiKey.Id, &apiKey.Name, &apiKey.Key, &apiKey.UserId, &apiKey.ProjectId, &apiKey.Revoked, &apiKey.CreatedAt)

	return apiKey, err
}