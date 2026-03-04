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

func (r Repository) GetLogsByProjectId(projectId string, startDate *string, endDate *string) ([]Log, error) {

	query := `
	SELECT path, latency, status, ip_address, timestamp FROM logs
	WHERE project_id = $1
	AND ($1::timestamp IS NULL OR created_at >= $1)
	AND ($2::timestamp IS NULL OR created_at <= $2)
	`

	rows, err := r.db.Query(context.Background(), query, projectId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	logs := []Log{}

	for rows.Next() {
		var log Log
		err := rows.Scan(&log.Path, &log.Latency, &log.Status, &log.IPAddress, &log.Timestamp)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}
