package log

import (
	"context"
	"fmt"
	"log"

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
	INSERT INTO logs (path, latency, status, ip_address, timestamp, api_key_id) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(context.Background(), query, log.Path, log.Latency, log.Status, log.IPAddress, log.Timestamp, log.ApiKey)

	return err
}

func (r Repository) GetLogsByProjectId(projectId string, startDate *string, endDate *string, statusCode *int, apiKeyType *string, page int, limit int) ([]LogResponse, error) {

	query := `
	SELECT logs.id, logs.path, logs.latency, logs.status, logs.ip_address, logs.timestamp, api_keys.name
	FROM logs
	JOIN api_keys ON logs.api_key_id = api_keys.id
	WHERE api_keys.project_id = $1
	  AND ($4::INTEGER IS NULL OR (logs.status >= $4 AND logs.status < $4 + 100))
	  AND ($2::TIMESTAMPTZ IS NULL OR logs.timestamp >= $2::TIMESTAMPTZ)
	  AND ($3::TIMESTAMPTZ IS NULL OR logs.timestamp <= $3::TIMESTAMPTZ)
	  AND ($5::TEXT IS NULL OR api_keys.name = $5)
	ORDER BY logs.timestamp DESC
	LIMIT $6 OFFSET $7
	`

	log.Printf("GetLogsByProjectId params: projectId=%s, startDate=%v, endDate=%v, statusCode=%v, apiKeyType=%v, limit=%d, offset=%d", projectId, startDate, endDate, statusCode, apiKeyType, limit, page*limit)

	rows, err := r.db.Query(context.Background(), query, projectId, startDate, endDate, statusCode, apiKeyType, limit, page*limit)

	if err != nil {
		log.Printf("Error getting logs: %v", err)
		return []LogResponse{}, err
	}

	defer rows.Close()

	logs := []LogResponse{}

	for rows.Next() {
		var log LogResponse
		err := rows.Scan(&log.Id, &log.Path, &log.Latency, &log.Status, &log.IPAddress, &log.Timestamp, &log.ApiKey)

		if err != nil {
			fmt.Println("error ", err)
			return []LogResponse{}, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func (r Repository) GetLogCountByProjectId(projectId string) int {

	var total int

	query := `
	SELECT COUNT(*) as total FROM logs
	JOIN api_keys ON logs.api_key_id = api_keys.id
	WHERE api_keys.project_id = $1
	`

	row := r.db.QueryRow(context.Background(), query, projectId)

	row.Scan(&total)

	return total

}
