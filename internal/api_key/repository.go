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

func (r Repository) FindById(id string) (ApiKey, error) {

	var apiKey ApiKey

	query := `
		SELECT id, name, key, user_id, project_id, revoked, created_at FROM api_keys WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&apiKey.Id, &apiKey.Name, &apiKey.Key, &apiKey.UserId, &apiKey.ProjectId, &apiKey.Revoked, &apiKey.CreatedAt)

	return apiKey, err
}

func (r Repository) FindAllByUserIdAndProjectId(userId string, projectId string) ([]ApiKey, error) {

	apiKeys := make([]ApiKey, 0)

	query := `
		SELECT id, name, key, created_at FROM api_keys WHERE revoked = FALSE AND user_id = $1 AND project_id = $2
	`

	rows, err := r.db.Query(context.Background(), query, userId, projectId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var apiKey ApiKey

		err := rows.Scan(&apiKey.Id, &apiKey.Name, &apiKey.Key, &apiKey.CreatedAt)

		if err != nil {
			return nil, err
		}

		apiKeys = append(apiKeys, apiKey)
	}

	return apiKeys, err
}

func (r Repository) RevokeApiKey(id string) error {

	query := `
		UPDATE api_keys SET revoked = true WHERE id = $1
	`
	_, err := r.db.Exec(context.Background(), query, id)

	return err
}
