package refreshtoken

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

func (r Repository) Save(refreshToken RefreshToken) error {

	query := `
		INSERT INTO refresh_tokens (id, token, expires_at, blacklisted, user_id) VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(context.Background(), query, refreshToken.Id, refreshToken.Token, refreshToken.ExpiresAt, refreshToken.Blacklisted, refreshToken.UserId)

	return err
}