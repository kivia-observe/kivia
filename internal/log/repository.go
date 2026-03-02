package log

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

func (r Repository) Save(log Log) error {

	query := `
	INSERT INTO logs (path, latency, status, ip_address, timestamp, project_Id) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(context.Background(), query, log.Path, log.Latency, log.Status, log.IPAddress, log.Timestamp, log.ProjectId)

	return err
}