package emailverification

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ev EmailVerification) error {
	query := `
		INSERT INTO email_verifications (id, email, otp_hash, attempts, expires_at, created_at, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(context.Background(), query,
		ev.Id, ev.Email, ev.OtpHash, ev.Attempts, ev.ExpiresAt, ev.CreatedAt, ev.UserId,
	)
	return err
}

func (r *Repository) FindByEmail(email string) EmailVerification {
	var ev EmailVerification
	query := `
		SELECT id, email, otp_hash, attempts, expires_at, created_at, user_id
		FROM email_verifications
		WHERE email = $1
		ORDER BY created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRow(context.Background(), query, email)
	row.Scan(&ev.Id, &ev.Email, &ev.OtpHash, &ev.Attempts, &ev.ExpiresAt, &ev.CreatedAt, &ev.UserId)
	return ev
}

func (r *Repository) IncrementAttempts(id string) error {
	query := `UPDATE email_verifications SET attempts = attempts + 1 WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *Repository) DeleteByUserId(userId string) error {
	query := `DELETE FROM email_verifications WHERE user_id = $1`
	_, err := r.db.Exec(context.Background(), query, userId)
	return err
}
