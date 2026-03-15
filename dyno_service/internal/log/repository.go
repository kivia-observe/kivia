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

func (r Repository) GetLogsByProjectId(projectId string, startDate *string, endDate *string, page int, limit int) ([]Log, error) {

	query := `
	SELECT id, path, latency, status, ip_address, timestamp
	FROM logs
	WHERE project_id = $1
	AND ($2::timestamp IS NULL OR timestamp >= $2)
	AND ($3::timestamp IS NULL OR timestamp <= $3)
	ORDER BY timestamp DESC
	LIMIT $4 OFFSET $5
	`

	rows, err := r.db.Query(context.Background(), query, projectId, startDate, endDate, limit, page * limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	logs := []Log{}

	for rows.Next() {
		var log Log
		err := rows.Scan(&log.Id, &log.Path, &log.Latency, &log.Status, &log.IPAddress, &log.Timestamp)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func (r Repository) GetLogCountByProjectId(projectId string) int {
	
	var total int

	query := `
		SELECT COUNT(*) as total FROM logs WHERE project_id = $1
	`

	row := r.db.QueryRow(context.Background(), query, projectId)

	row.Scan(&total)

	return total

}
